package auth

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/gin-gonic/gin"
)

type User struct {
	ui application.UserAppInterface
}

func NewUser(ui application.UserAppInterface) *User {
	return &User{
		ui: ui,
	}
}

func (u *User) Login(c *gin.Context) {

}

func (u *User) Register(c *gin.Context) {

}
