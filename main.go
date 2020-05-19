package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/lihanx-blog/diyBlog/routers"
)

func main() {
	router := gin.Default()
	router.GET("/", routers.Index)
}