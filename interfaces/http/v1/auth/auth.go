package auth

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserRouter struct {
	ui application.UserAppInterface
}

func NewUserRouter(ui application.UserAppInterface) *UserRouter {
	return &UserRouter{
		ui: ui,
	}
}
