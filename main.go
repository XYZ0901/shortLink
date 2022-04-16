package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shortLink/api"
	"shortLink/middleware"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())

	{
		r.POST("/api/shorten", api.Shorten)
		r.GET("/s/:short_link", api.Redirect)
	}

	s := http.Server{
		Addr:    ":8999",
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("server listenAndServer failed: ", err)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln("server shutdown failed: ", err)
	}
	log.Println("Server exiting")
}
