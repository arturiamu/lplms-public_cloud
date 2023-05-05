package auth

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/arturiamu/lplms-public_cloud/common"
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/utils/ctx"
	"github.com/arturiamu/lplms-public_cloud/utils/md5util"
	"github.com/arturiamu/lplms-public_cloud/utils/token"
	"github.com/arturiamu/lplms-public_cloud/utils/uuid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

var validate = validator.New()

type User struct {
	ui application.UserAppInterface
}

func NewUser(ui application.UserAppInterface) *User {
	return &User{
		ui: ui,
	}
}

type RegisterArgs struct {
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required,gte=6"`
	RePassword string     `json:"re_password" validate:"eqfield=Password"`
	Email      string     `json:"email" validate:"required,email"`
	Telephone  string     `json:"telephone" validate:"omitempty"`
	Birthday   *time.Time `json:"birthday" validate:"omitempty,datetime"`
}

func (u *User) Register(c *gin.Context) {
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
	_, err := u.ui.SaveUser(&entity.UserCreateArg{User: user})
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

func (u *User) Login(c *gin.Context) {
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
	resp, err := u.ui.GetUserBy(&entity.UserGetByArgs{Email: args.Email})
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

func (u *User) Logout(c *gin.Context) {
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

func (u *User) UpdateInfo(c *gin.Context) {
	var (
		args UpdateInfoArgs
		user = ctx.GetUser(c)
	)
	if err := c.BindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	if err := validate.Struct(&args); err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
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

	_, err := u.ui.UpdateUser(&entity.UserUpdateArg{User: *user})
	if err != nil {
		c.JSON(http.StatusBadRequest, common.FailWith(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, common.Success())
}
