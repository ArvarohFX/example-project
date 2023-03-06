package dao

import (
	"time"
)

type User struct {
	ID        string  `gorm:"column:id"`
	FirstName *string `gorm:"column:first_name"`
	LastName  *string `gorm:"column:last_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users.users"
}
