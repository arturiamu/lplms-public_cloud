package auth

import (
	"github.com/arturiamu/lplms-public_cloud/application"
	"github.com/go-playground/validator/v10"
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
