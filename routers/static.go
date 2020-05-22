package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

func init() {
	register("static", func(router *gin.Engine) {
		staticBox := packr.New("static", "../static")
		router.StaticFS("/static", staticBox)
	})
}