package persistence

import (
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
	"github.com/jinzhu/gorm"
)

var _ repository.UserRepository = &UserRepo{}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
