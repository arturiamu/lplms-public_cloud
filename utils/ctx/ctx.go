package ctx

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/interfaces/middleware"
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
	user, ok := u.(entity.User)
	if ok {
		return &user
	}
	return nil
}
