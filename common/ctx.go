package common

import (
	"github.com/gin-gonic/gin"
)

// GetUser
// 只包含 uid 和 project_id 信息
//func GetUser(c *gin.Context) *token.AuthUser {
//	if c == nil {
//		return nil
//	}
//	u, ok := c.Get(UserInfo)
//	if !ok {
//		return nil
//	}
//	user, ok := u.(token.AuthUser)
//	if ok {
//		return &user
//		//return &entity.User{
//		//	UID:       user.UID,
//		//	ProjectID: user.ProjectID,
//		//}
//	}
//	return nil
//}

func GetUid(c *gin.Context) string {
	ctxUid, ok := c.Get(UserID)
	if !ok {
		return ""
	}
	uid, ok := ctxUid.(string)
	if !ok {
		return ""
	}
	return uid
}

func GetProject(c *gin.Context) string {
	ctxPid, ok := c.Get(UserID)
	if !ok {
		return ""
	}
	pid, ok := ctxPid.(string)
	if !ok {
		return ""
	}
	return pid
}
