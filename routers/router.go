package routers

import (
	"errors"
	"sync"
	"fmt"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lihanx/gin-blog/utils"
	"github.com/lihanx/gin-blog/ginengine"
)

type routerFunc struct {
	Name	string
	Weight	int
	Func 	func(router *gin.Engine)
}

type routerSlice []routerFunc

func (r routerSlice) Len() int { return len(r) }

func (r routerSlice) Less(i, j int) bool { return r[i].Weight > r[j].Weight }

func (r routerSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i]}

var userRouterOnce sync.Once
var routers routerSlice

func register(name string, f func(router *gin.Engine)) {
	registerWithWeight(name, 50, f)
}

// 注册路由
func registerWithWeight(name string, weight int, f func(router *gin.Engine)) {
	if weight > 100 || weight < 0 {
		utils.CheckAndExit(errors.New(fmt.Sprintf("router weight must be >= 0 and <= 100")))
	}
	for _, r := range routers {
		if strings.ToLower(name) == strings.ToLower(r.Name) {
			utils.CheckAndExit(errors.New(fmt.Sprintf("router [%s] has already been registered", r.Name)))
		}
	}

	routers  = append(routers, routerFunc{
		Name:	name,
		Weight:	weight,
		Func:	f,
	})
}

func Setup() {
	userRouterOnce.Do(func() {
		sort.Sort(routers)
		for _, r := range routers {
			r.Func(ginengine.Engine)
		}
	})
}