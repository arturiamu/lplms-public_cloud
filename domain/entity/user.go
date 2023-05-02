package entity

import (
	"time"
)

type User struct {
	UID       string     `gorm:"column:uid;primary_key"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password;type:char(32);not null"`
	Email     string     `gorm:"column:email;unique"`
	ProjectID string     `gorm:"column:project_id;unique"`
	Telephone string     `gorm:"column:telephone;unique"`
	Birthday  *time.Time `gorm:"column:birthday"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time  `gorm:"column:deleted_at;index"`
	Role      string     `gorm:"column:role;default:ordinary"`
	Project   string     `gorm:"column:project"`
}

func (u User) TableName() string {
	return "users"
}

type UserCreateArg struct {
}

type UserDeleteArg struct {
}

type UserUpdateArg struct {
}

type UserGetArg struct {
}

type UserListArg struct {
}

type UserCreateResp struct{}

type UserDeleteResp struct{}

type UserUpdateResp struct{}

type UserGetResp struct{}

type UserListResp struct{}
