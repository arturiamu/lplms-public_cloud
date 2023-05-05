package entity

import (
	"time"
)

type User struct {
	UID       string     `gorm:"column:uid;primary_key;not null" json:"uid"`
	Username  string     `gorm:"column:username;not null" json:"username"`
	Password  string     `gorm:"column:password;type:char(32);not null" json:"password"`
	Email     string     `gorm:"column:email;unique;not null" json:"email"`
	ProjectID string     `gorm:"column:project_id;unique;not null" json:"project_id"`
	Telephone string     `gorm:"column:telephone;unique" json:"telephone"`
	Birthday  *time.Time `gorm:"column:birthday" json:"birthday"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time  `gorm:"column:deleted_at;index" json:"deleted_at"`
	Role      string     `gorm:"column:role;default:ordinary" json:"role"`
}

func (u User) TableName() string {
	return "users"
}

type UserCreateArg struct {
	User User
}

type UserDeleteArg struct {
	UID string
}

type UserUpdateArg struct {
	User User
}

type UserGetArg struct {
	UID string
}

type UserListArg struct {
	UID string
}

type UserCreateResp struct {
	User User
}

type UserDeleteResp struct {
	User User
}

type UserUpdateResp struct {
	User User
}

type UserGetResp struct {
	User User
}

type UserListResp struct {
	Users []User
}

type UserGetByArgs struct {
	Email string
}

type UserGetByResp struct {
	User User
}
