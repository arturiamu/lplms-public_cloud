package auth

import (
	"fmt"
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/ctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ui application.UserAppInterface
}

func NewUser(ui application.UserAppInterface) *User {
	return &User{
		ui: ui,
	}
}

func (u *User) Register(c *gin.Context) {

}

func (u *User) Login(c *gin.Context) {
	u.ui.SaveUser(&entity.User{})
}

func (u *User) Logout(c *gin.Context) {
	// TODO clear redis、session
	c.JSON(http.StatusOK, "OK")
}

func (u *User) UpdateInfo(c *gin.Context) {
	var (
		user = ctx.GetUser(c)
	)
	fmt.Println(user)

}
