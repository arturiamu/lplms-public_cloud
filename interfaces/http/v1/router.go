package v1

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/auth"
	"github.com/arturiamu/lplms-public_cloud/interfaces/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(stack application.StackInterface, engine *gin.Engine) {
	engine.Use(middleware.CORSMiddleware(), middleware.AuthMiddleware())
	v1 := engine.Group("lplms/api/v1")
	RegisterUserGroup(v1, stack.User())
}

func RegisterUserGroup(rg *gin.RouterGroup, ui application.UserAppInterface) {
	user := auth.NewUser(ui)
	rg.POST("/user", user.Register)
}
