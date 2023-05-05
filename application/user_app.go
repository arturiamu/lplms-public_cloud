package application

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
)

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(args *entity.UserCreateArg) (*entity.UserCreateResp, error)
	DeleteUser(args *entity.UserDeleteArg) (*entity.UserDeleteResp, error)
	UpdateUser(args *entity.UserUpdateArg) (*entity.UserUpdateResp, error)
	GetUser(args *entity.UserGetArg) (*entity.UserGetResp, error)
	ListUser(args *entity.UserListArg) (*entity.UserListResp, error)
	GetUserBy(args *entity.UserGetByArgs) (*entity.UserGetByResp, error)
}

type userApp struct {
	us repository.UserRepositoryInterface
}

func (u *userApp) DeleteUser(args *entity.UserDeleteArg) (*entity.UserDeleteResp, error) {
	return u.us.DeleteUser(args)
}

func (u *userApp) UpdateUser(args *entity.UserUpdateArg) (*entity.UserUpdateResp, error) {
	return u.us.UpdateUser(args)
}

func (u *userApp) SaveUser(args *entity.UserCreateArg) (*entity.UserCreateResp, error) {
	return u.us.SaveUser(args)
}

func (u *userApp) GetUser(args *entity.UserGetArg) (*entity.UserGetResp, error) {
	return u.us.GetUser(args)
}

func (u *userApp) ListUser(args *entity.UserListArg) (*entity.UserListResp, error) {
	return u.us.ListUser(args)
}

func (u *userApp) GetUserBy(args *entity.UserGetByArgs) (*entity.UserGetByResp, error) {
	return u.us.GetUserBy(args)
}
