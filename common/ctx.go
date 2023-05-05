package common

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/interfaces/middleware"
	"github.com/arturiamu/lplms-public_cloud/utils/token"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) *entity.User {
	if c == nil {
		return nil
	}
	u, ok := c.Get(middleware.UserInfo)
	if !ok {
		return nil
	}
	user, ok := u.(token.AuthUser)
	if ok {
		return &entity.User{
			UID:       user.UID,
			ProjectID: user.ProjectID,
		}
	}
	return nil
}

func GetProject(c *gin.Context) string {
	user := GetUser(c)
	if user == nil {
		return ""
	}
	return user.ProjectID
}
