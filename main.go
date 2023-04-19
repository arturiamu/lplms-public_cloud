package main

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	v1 "github.com/arturiamu/lplms-public_cloud/interfaces/http/v1"
	"github.com/gin-gonic/gin"
)

func main() {

	stack := application.NewStack(nil, nil)
	engine := gin.New()
	v1.InitRouter(stack, engine)
}
