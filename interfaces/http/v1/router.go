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
	user := auth.NewUser(ui)
	rg.POST("/register", user.Register)
	rg.POST("/login", user.Login)
	rg.PUT("/info", middleware.AuthMiddleware(), user.UpdateInfo)
	rg.GET("/logout", middleware.AuthMiddleware(), user.Logout)
}

func RegisterComputeGroup(rg *gin.RouterGroup, ci application.ComputeAppInterface) {
	rg.Any("/ping", ping)
	c := compute.NewCompute(ci)
	rg.POST("/server", c.CreateServer)
	rg.DELETE("/server/:id", c.DeleteServer)
	rg.PATCH("/server/:id", c.UpdateServer)
	rg.GET("/server/:id", c.GetServer)
	rg.GET("/server", c.ListServer)
	rg.GET("/server/:id/disks", c.GetServerDisks)

	rg.POST("/flavor", c.CreateFlavor)
	rg.DELETE("/flavor/:id", c.DeleteFlavor)
	rg.PATCH("/flavor/:id", c.UpdateFlavor)
	rg.GET("/flavor/:id", c.GetFlavor)
	rg.GET("/flavor", c.ListFlavor)

	rg.POST("/image", c.CreateImage)
	rg.DELETE("/image/:id", c.DeleteImage)
	rg.PATCH("/image/:id", c.UpdateImage)
	rg.GET("/image/:id", c.GetImage)
	rg.GET("/image", c.ListImage)

	rg.POST("/keypair", c.CreateKeypair)
	rg.DELETE("/keypair/:id", c.DeleteKeypair)
	rg.PATCH("/keypair/:id", c.UpdateKeypair)
	rg.GET("/keypair/:id", c.GetKeypair)
	rg.GET("/keypair", c.ListKeypair)
	rg.GET("/keypair/:id/attach", c.AttachKeyPair)
	rg.GET("/keypair/:id/detach", c.DetachKeyPair)

	rg.POST("/security_group", c.CreateSecurityGroup)
	rg.DELETE("/security_group/:id", c.DeleteSecurityGroup)
	rg.PATCH("/security_group/:id", c.UpdateSecurityGroup)
	rg.GET("/security_group/:id", c.GetSecurityGroup)
	rg.GET("/security_group", c.ListSecurityGroup)

	rg.POST("/security_group/:id/role", c.CreateSecurityGroupRule)
	rg.DELETE("/security_group/:id/role/:id", c.DeleteSecurityGroupRule)
	rg.GET("/security_group/:id/role/:id", c.GetSecurityGroupRule)
	rg.GET("/security_group/:id/role", c.ListSecurityGroupRule)
}

func RegisterStorageGroup(rg *gin.RouterGroup, si application.StorageAppInterface) {
	rg.Any("/ping", ping)
	s := storage.NewStorage(si)
	rg.POST("/disk", s.CreateDisk)
	rg.DELETE("/disk/:id", s.DeleteDisk)
	rg.PATCH("/disk/:id", s.UpdateDisk)
	rg.GET("/disk/:id", s.GetDisk)
	rg.GET("/disk", s.ListDisk)

	rg.POST("/disk/:id/attach", s.AttachDisk)
	rg.POST("/disk/:id/detach", s.DetachDisk)
	rg.POST("/disk/:id/resize", s.ResizeDisk)
	rg.POST("/disk/:id/reset", s.ResetDisk)

	rg.POST("/snapshot", s.CreateSnapshot)
	rg.DELETE("/snapshot/:id", s.DeleteSnapshot)
	rg.PATCH("/snapshot/:id", s.UpdateSnapshot)
	rg.GET("/snapshot/:id", s.GetSnapshot)
	rg.GET("/snapshot", s.ListSnapshot)
}

func RegisterNetworkGroup(rg *gin.RouterGroup, ni application.NetworkAppInterface) {
	rg.Any("/ping", ping)
	n := network.NewNetwork(ni)
	rg.POST("/eip", n.CreateEip)
	rg.DELETE("/eip/:id", n.DeleteEip)
	rg.PATCH("/eip/:id", n.UpdateEip)
	rg.GET("/eip/:id", n.GetEip)
	rg.GET("/eip", n.ListEip)

	rg.POST("/vpc", n.CreateVpc)
	rg.DELETE("/vpc/:id", n.DeleteVpc)
	rg.PATCH("/vpc/:id", n.UpdateVpc)
	rg.GET("/vpc/:id", n.GetVpc)
	rg.GET("/vpc", n.ListVpc)

	//rg.POST("/slb", n.CreateSlb)
	//rg.DELETE("/slb/:id", n.DeleteSlb)
	//rg.PATCH("/slb/:id", n.UpdateSlb)
	//rg.GET("/slb/:id", n.GetSlb)
	//rg.GET("/slb", n.ListSlb)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func noRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"err": "no route"})
}
