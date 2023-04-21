package middleware

import (
	"github.com/arturiamu/lplms-public_cloud/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

const UserInfo = "user_info"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if len(tokenStr) == 0 {
			c.AbortWithError(http.StatusNonAuthoritativeInfo, http.ErrNoCookie)
			return
		}
		ac, err := token.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithError(http.StatusNonAuthoritativeInfo, err)
			return
		}
		c.Set(UserInfo, ac.User)
		c.Next()
		return
	}
}
