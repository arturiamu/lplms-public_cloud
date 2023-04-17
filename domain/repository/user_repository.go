package repository

import "github.com/arturiamu/lplms-public_cloud/domain/entity"

type UserRepositoryInterface interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	ListUser() (*entity.User, error)
	GetUserBy(uint64) (*entity.User, error)
}
