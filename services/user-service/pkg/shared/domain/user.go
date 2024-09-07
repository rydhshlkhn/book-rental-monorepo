package domain

import (
	"time"
)

// User model
type User struct {
	ID           int       `gorm:"column:id;primary_key" json:"id"`
	Field        string    `gorm:"column:field;type:varchar(255)" json:"field"`
	Username     string    `gorm:"column:username;type:varchar(255)" json:"username"`
	Email        string    `gorm:"column:email;type:varchar(255)" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255)" json:"password_hash"`
	PasswordSalt string    `gorm:"column:password_salt;type:varchar(255)" json:"password_salt"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of User model
func (User) TableName() string {
	return "users"
}
