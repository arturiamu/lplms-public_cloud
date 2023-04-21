package v1

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/auth"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/compute"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/network"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/storage"
	"github.com/arturiamu/lplms-public_cloud/interfaces/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(stack application.StackInterface, engine *gin.Engine) {
	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(middleware.CORSMiddleware(), middleware.AuthMiddleware())

	v1 := engine.Group("lplms/api/v1", middleware.CORSMiddleware())
	RegisterUserGroup(v1.Group("/auth"), stack.User())
	RegisterComputeGroup(v1.Group("/compute", middleware.AuthMiddleware()), stack.Compute())
	RegisterStorageGroup(v1.Group("/storage", middleware.AuthMiddleware()), stack.Storage())
	RegisterNetworkGroup(v1.Group("/network", middleware.AuthMiddleware()), stack.Network())
}

func RegisterUserGroup(rg *gin.RouterGroup, ui application.UserAppInterface) {
	user := auth.NewUser(ui)
	rg.POST("/register", user.Register)
	rg.POST("/login", user.Login)
	rg.PUT("/user", middleware.AuthMiddleware(), user.UpdateInfo)
	rg.GET("/logout", user.Logout)
}

func RegisterComputeGroup(rg *gin.RouterGroup, ci application.ComputeAppInterface) {
	c := compute.NewCompute(ci)
	rg.POST("/server", c.CreateServer)
	rg.DELETE("/server/:id", c.DeleteServer)
}

func RegisterStorageGroup(rg *gin.RouterGroup, si application.StorageAppInterface) {
	s := storage.NewStorage(si)
	rg.POST("/storage", s.CreateDisk)
}

func RegisterNetworkGroup(rg *gin.RouterGroup, ni application.NetworkAppInterface) {
	n := network.NewNetwork(ni)
	rg.POST("/network", n.CreateVpc)
}
