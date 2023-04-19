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

	v1 := engine.Group("lplms/api/v1")
	RegisterUserGroup(v1, stack.User())
	RegisterComputeGroup(v1, stack.Compute())
	RegisterStorageGroup(v1, stack.Storage())
	RegisterNetworkGroup(v1, stack.Network())
}

func RegisterUserGroup(rg *gin.RouterGroup, ui application.UserAppInterface) {
	user := auth.NewUser(ui)
	rg.POST("/user", user.Register)
}

func RegisterComputeGroup(rg *gin.RouterGroup, ci application.ComputeAppInterface) {
	c := compute.NewCompute(ci)
	rg.POST("/server", c.CreateServer)
	rg.DELETE("/server/:id", c.DeleteServer)
}

func RegisterStorageGroup(rg *gin.RouterGroup, si application.StorageAppInterface) {
	s := storage.NewStorage(si)
	rg.POST("/server", s.CreateDisk)
}

func RegisterNetworkGroup(rg *gin.RouterGroup, ni application.NetworkAppInterface) {
	n := network.NewNetwork(ni)
	rg.POST("/server", n.CreateVpc)
}
