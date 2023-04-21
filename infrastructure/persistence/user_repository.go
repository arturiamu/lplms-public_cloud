package persistence

import (
	"github.com/arturiamu/lplms-public_cloud/domain/entity"
	"github.com/arturiamu/lplms-public_cloud/domain/repository"
	"gorm.io/gorm"
)

var _ repository.UserRepositoryInterface = &UserRepo{}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	var userRepo = UserRepo{db: db}
	stk.U = &userRepo
	return &userRepo
}

func (u *UserRepo) SaveUser(user *entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetUser(u2 uint64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) ListUser() (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetUserBy(u2 uint64) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
