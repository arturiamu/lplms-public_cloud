package v1

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/auth"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/compute"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/network"
	"github.com/arturiamu/lplms-public_cloud/interfaces/http/v1/storage"
	"github.com/arturiamu/lplms-public_cloud/interfaces/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(stack application.StackInterface, engine *gin.Engine) {
	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(middleware.CORSMiddleware())
	engine.NoRoute(noRoute)
	v1 := engine.Group("/lplms/api/v1")
	v1.Any("/ping", ping)
	RegisterUserGroup(v1.Group("/auth"), stack.User())
	RegisterComputeGroup(v1.Group("/compute", middleware.AuthMiddleware()), stack.Compute())
	RegisterStorageGroup(v1.Group("/storage", middleware.AuthMiddleware()), stack.Storage())
	RegisterNetworkGroup(v1.Group("/network", middleware.AuthMiddleware()), stack.Network())
}

func RegisterUserGroup(rg *gin.RouterGroup, ui application.UserAppInterface) {
	rg.Any("/ping", ping)
	userRouter := auth.NewUserRouter(ui)
	rg.POST("/register", userRouter.Register)
	rg.POST("/login", userRouter.Login)
	rg.PUT("/info", middleware.AuthMiddleware(), userRouter.UpdateInfo)
	rg.GET("/logout", middleware.AuthMiddleware(), userRouter.Logout)
}

func RegisterComputeGroup(rg *gin.RouterGroup, ci application.ComputeAppInterface) {
	rg.Any("/ping", ping)
	routerCompute := compute.NewCompute(ci)
	rg.POST("/server", routerCompute.CreateServer)
	rg.DELETE("/server/:id", routerCompute.DeleteServer)
	rg.PATCH("/server/:id", routerCompute.UpdateServer)
	rg.GET("/server/:id", routerCompute.GetServer)
	rg.GET("/server", routerCompute.ListServer)
	rg.GET("/server/:id/disks", routerCompute.GetServerDisks)
	rg.POST("/server/:id/start", routerCompute.StartServer)
	rg.POST("/server/:id/stop", routerCompute.StopServer)
	rg.POST("/server/:id/restart", routerCompute.RestartServer)

	rg.POST("/flavor", routerCompute.CreateFlavor)
	rg.DELETE("/flavor/:id", routerCompute.DeleteFlavor)
	rg.PATCH("/flavor/:id", routerCompute.UpdateFlavor)
	rg.GET("/flavor/:id", routerCompute.GetFlavor)
	rg.GET("/flavor", routerCompute.ListFlavor)

	rg.POST("/image", routerCompute.CreateImage)
	rg.DELETE("/image/:id", routerCompute.DeleteImage)
	rg.PATCH("/image/:id", routerCompute.UpdateImage)
	rg.GET("/image/:id", routerCompute.GetImage)
	rg.GET("/image", routerCompute.ListImage)

	rg.POST("/keypair", routerCompute.CreateKeypair)
	rg.DELETE("/keypair/:id", routerCompute.DeleteKeypair)
	rg.PATCH("/keypair/:id", routerCompute.UpdateKeypair)
	rg.GET("/keypair/:id", routerCompute.GetKeypair)
	rg.GET("/keypair", routerCompute.ListKeypair)
	rg.GET("/keypair/:id/attach", routerCompute.AttachKeyPair)
	rg.GET("/keypair/:id/detach", routerCompute.DetachKeyPair)

	rg.POST("/security_group", routerCompute.CreateSecurityGroup)
	rg.DELETE("/security_group/:id", routerCompute.DeleteSecurityGroup)
	rg.PATCH("/security_group/:id", routerCompute.UpdateSecurityGroup)
	rg.GET("/security_group/:id", routerCompute.GetSecurityGroup)
	rg.GET("/security_group", routerCompute.ListSecurityGroup)

	rg.POST("/security_group/:id/role", routerCompute.CreateSecurityGroupRule)
	rg.DELETE("/security_group/:id/role/:id", routerCompute.DeleteSecurityGroupRule)
	rg.GET("/security_group/:id/role/:id", routerCompute.GetSecurityGroupRule)
	rg.GET("/security_group/:id/role", routerCompute.ListSecurityGroupRule)
}

func RegisterStorageGroup(rg *gin.RouterGroup, si application.StorageAppInterface) {
	rg.Any("/ping", ping)
	routerStorage := storage.NewStorage(si)
	rg.POST("/disk", routerStorage.CreateDisk)
	rg.DELETE("/disk/:id", routerStorage.DeleteDisk)
	rg.PATCH("/disk/:id", routerStorage.UpdateDisk)
	rg.GET("/disk/:id", routerStorage.GetDisk)
	rg.GET("/disk", routerStorage.ListDisk)

	rg.POST("/disk/:id/attach", routerStorage.AttachDisk)
	rg.POST("/disk/:id/detach", routerStorage.DetachDisk)
	rg.POST("/disk/:id/resize", routerStorage.ResizeDisk)
	rg.POST("/disk/:id/reset", routerStorage.ResetDisk)

	rg.POST("/snapshot", routerStorage.CreateSnapshot)
	rg.DELETE("/snapshot/:id", routerStorage.DeleteSnapshot)
	rg.PATCH("/snapshot/:id", routerStorage.UpdateSnapshot)
	rg.GET("/snapshot/:id", routerStorage.GetSnapshot)
	rg.GET("/snapshot", routerStorage.ListSnapshot)
}

func RegisterNetworkGroup(rg *gin.RouterGroup, ni application.NetworkAppInterface) {
	rg.Any("/ping", ping)
	routerNetwork := network.NewNetwork(ni)
	rg.POST("/eip", routerNetwork.CreateEip)
	rg.DELETE("/eip/:id", routerNetwork.DeleteEip)
	rg.PATCH("/eip/:id", routerNetwork.UpdateEip)
	rg.GET("/eip/:id", routerNetwork.GetEip)
	rg.GET("/eip", routerNetwork.ListEip)

	rg.POST("/vpc", routerNetwork.CreateVpc)
	rg.DELETE("/vpc/:id", routerNetwork.DeleteVpc)
	rg.PATCH("/vpc/:id", routerNetwork.UpdateVpc)
	rg.GET("/vpc/:id", routerNetwork.GetVpc)
	rg.GET("/vpc", routerNetwork.ListVpc)

	rg.POST("/v_switch", routerNetwork.CreateVSwitch)
	rg.DELETE("/v_switch/:id", routerNetwork.DeleteVSwitch)
	rg.PATCH("/v_switch/:id", routerNetwork.UpdateVSwitch)
	rg.GET("/v_switch/:id", routerNetwork.GetVSwitch)
	rg.GET("/v_switch", routerNetwork.ListVSwitch)

	//rg.POST("/slb", routerNetwork.CreateSlb)
	//rg.DELETE("/slb/:id", routerNetwork.DeleteSlb)
	//rg.PATCH("/slb/:id", routerNetwork.UpdateSlb)
	//rg.GET("/slb/:id", routerNetwork.GetSlb)
	//rg.GET("/slb", routerNetwork.ListSlb)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func noRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"err": "no route"})
}
