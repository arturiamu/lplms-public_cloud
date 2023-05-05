package persistence

import (
	"errors"
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

func (u *UserRepo) DeleteUser(args *entity.UserDeleteArg) (*entity.UserDeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) UpdateUser(args *entity.UserUpdateArg) (*entity.UserUpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) SaveUser(args *entity.UserCreateArg) (*entity.UserCreateResp, error) {
	err := u.db.Save(&args.User).Error
	if err != nil {
		return nil, err
	}
	return nil, err
}

// GetUser
// get user by uid
func (u *UserRepo) GetUser(args *entity.UserGetArg) (*entity.UserGetResp, error) {
	var user entity.User
	err := u.db.Where("uid=?", args.UID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &entity.UserGetResp{User: user}, err
}

func (u *UserRepo) ListUser(args *entity.UserListArg) (*entity.UserListResp, error) {
	var users []entity.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &entity.UserListResp{Users: users}, err
}

// GetUserBy
// get user by email
func (u *UserRepo) GetUserBy(args *entity.UserGetByArgs) (*entity.UserGetByResp, error) {
	var user entity.User
	err := u.db.Where("email=?", args.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("username or password err")
		}
		return nil, err
	}
	return &entity.UserGetByResp{User: user}, err
}

///////////////////////////// help functions /////////////
