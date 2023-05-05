package auth

import (
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/md5util"
	"github.com/arturiamu/lplms-public_cloud/utils/token"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RegisterArgs struct {
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required,gte=6"`
	RePassword string     `json:"re_password" validate:"eqfield=Password"`
	Email      string     `json:"email" validate:"required,email"`
	Telephone  string     `json:"telephone" validate:"omitempty"`
	Birthday   *time.Time `json:"birthday" validate:"omitempty,datetime"`
}

// Register 注册服务
func (ur *UserRouter) Register(c *gin.Context) {
	var (
		args RegisterArgs
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	if err := validate.Struct(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	user := entity.User{
		UID:       md5util.String2MD5Str(uuid.GenerateUUID()),
		Username:  args.Username,
		Password:  md5util.String2MD5Str(args.Password),
		ProjectID: md5util.String2MD5Str(args.Email),
		Email:     args.Email,
		Telephone: args.Telephone,
		Birthday:  args.Birthday,
	}
	_, err := ur.ui.SaveUser(&entity.UserCreateArg{User: user})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}

type LoginArgs struct {
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,gte=6"`
	VerificationCode string `json:"verification_code"`
}

// Login 登录服务
func (ur *UserRouter) Login(c *gin.Context) {
	var (
		args LoginArgs
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	if err := validate.Struct(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	resp, err := ur.ui.GetUserBy(&entity.UserGetByArgs{Email: args.Email})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	if resp.User.Password != md5util.String2MD5Str(args.Password) {
		c.JSON(http.StatusBadRequest, common.FailWith("username or password err", nil))
		return
	}
	tk, err := token.GenerateToken(&resp.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.SuccessWith("", tk))
}

// Logout 登出
func (ur *UserRouter) Logout(c *gin.Context) {
	// TODO clear redis、session
	c.JSON(http.StatusOK, common.Success())
}

type UpdateInfoArgs struct {
	Username   *string    `json:"username" validate:"omitempty"`
	Password   *string    `json:"password" validate:"omitempty,gte=6"`
	RePassword *string    `json:"re_password" validate:"eqfield=Password"`
	Email      *string    `json:"email" validate:"omitempty,email"`
	Telephone  *string    `json:"telephone" validate:"omitempty"`
	Birthday   *time.Time `json:"birthday" validate:"omitempty,datetime"`
}

// UpdateInfo 修改用户信息
func (ur *UserRouter) UpdateInfo(c *gin.Context) {
	var (
		args UpdateInfoArgs
		u    = common.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	if err := validate.Struct(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	resp, err := ur.ui.GetUser(&entity.UserGetArg{UID: u.UID})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	user := resp.User
	if args.Username != nil {
		user.Username = *args.Username
	}
	if args.Password != nil {
		user.Password = *args.Password
	}
	if args.Email != nil {
		user.Email = *args.Email
	}
	if args.Telephone != nil {
		user.Telephone = *args.Telephone
	}
	if args.Birthday != nil {
		user.Birthday = args.Birthday
	}

	_, err = ur.ui.UpdateUser(&entity.UserUpdateArg{User: user})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}
