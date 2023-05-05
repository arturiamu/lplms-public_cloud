package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type UserRepositoryInterface interface {
	SaveUser(args *entity.UserCreateArg) (*entity.UserCreateResp, error)
	DeleteUser(args *entity.UserDeleteArg) (*entity.UserDeleteResp, error)
	UpdateUser(args *entity.UserUpdateArg) (*entity.UserUpdateResp, error)
	GetUser(args *entity.UserGetArg) (*entity.UserGetResp, error)
	ListUser(args *entity.UserListArg) (*entity.UserListResp, error)
	GetUserBy(args *entity.UserGetByArgs) (*entity.UserGetByResp, error)
}
