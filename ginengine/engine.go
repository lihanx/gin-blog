package ginengine

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine
var engineOnce sync.Once

func Init() {
	engineOnce.Do(func() {
		Engine = gin.New()
	})
}