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
