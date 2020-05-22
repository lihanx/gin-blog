package routers

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	register("web", func(router *gin.Engine) {
		router.GET("", index)
	})
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", "biubiubiu")
}