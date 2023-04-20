package main

import (
	"flag"
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/config"
	v1 "github.com/arturiamu/lplms-public_cloud/interfaces/http/v1"
	"github.com/gin-gonic/gin"
)

var confPath *string
var kubePath *string

func main() {
	flag.StringVar(confPath, "c", "", "app -f=/path/to/config.yaml")
	flag.StringVar(kubePath, "k", "config/config", "app -f=/path/to/config")
	if err := config.LoadConfig(*confPath); err != nil {
		panic(err)
	}

	//stack := application.NewStack(config.AppConfig, nil)
	stack, err := application.NewStack(kubePath)
	if err != nil {
		panic(err)
	}
	engine := gin.New()
	v1.InitRouter(stack, engine)
}
