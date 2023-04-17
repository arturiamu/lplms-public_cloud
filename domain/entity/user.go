package entity

type User struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
}
