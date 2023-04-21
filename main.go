package main

import (
	"flag"
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/config"
	v1 "github.com/arturiamu/lplms-public_cloud/interfaces/http/v1"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var confPath string
var kubePath string

func main() {
	flag.StringVar(&confPath, "c", "", "app -f=/path/to/config.yaml")
	flag.StringVar(&kubePath, "k", "config/config", "app -f=/path/to/config")
	if err := config.LoadConfig(confPath); err != nil {
		panic(err)
	}

	stack, err := application.NewStack(&kubePath)
	if err != nil {
		panic(err)
	}
	engine := gin.New()
	v1.InitRouter(stack, engine)

	portStr := fmt.Sprintf("127.0.0.1:%d", config.AppConfig.App.Port)
	server := http.Server{
		Addr:    portStr,
		Handler: engine,
	}

	go func() {
		log.Println("try to start server at " + portStr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Println("get signal: ", sig)
	log.Println("ready to quit:")
	cleanup()
	log.Println("bye")
}

func cleanup() {
	// TODO
}
