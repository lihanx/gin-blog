package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	register("web", func(router *gin.Engine) {
		router.GET("", index)
		router.GET("/about", about)
		router.GET("/contact", contact)
	})
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"author": "李瀚翾"})
}

func about(c *gin.Context) {
	c.HTML(http.StatusOK, "about.tmpl", gin.H{"author": "李瀚翾"})
}

func contact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.tmpl", gin.H{"author": "李瀚翾"})
}