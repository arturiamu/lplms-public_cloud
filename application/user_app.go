package application

import (
	"lplms-public_cloud/domain/entity"
	"lplms-public_cloud/domain/repository"
)

var _ UserAppInterface = &userApp{}

type UserAppInterface interface {
	SaveUser(user *entity.User) (*entity.User, error)
	GetUser(id uint64) (*entity.User, error)
	ListUser() (*entity.User, error)
	GetUserBy(id uint64) (*entity.User, error)
}

type userApp struct {
	us repository.UserRepositoryInterface
}

func (u userApp) SaveUser(user *entity.User) (*entity.User, error) {
	return u.us.SaveUser(user)
}

func (u userApp) GetUser(id uint64) (*entity.User, error) {
	return u.us.GetUser(id)
}

func (u userApp) ListUser() (*entity.User, error) {
	return u.us.ListUser()
}

func (u userApp) GetUserBy(id uint64) (*entity.User, error) {
	return u.us.GetUserBy(id)
}
