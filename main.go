package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/gin-gonic/gin"

	"github.com/lihanx/gin-blog/routers"
	"github.com/lihanx/gin-blog/ginengine"
)

func main() {
	ginengine.Init()
	routers.Setup()

	srv := &http.Server{
		Addr:		"localhost:8080",
		Handler:	ginengine.Engine,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Println("ttt")
}