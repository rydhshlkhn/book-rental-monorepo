package domain

import (
	"time"
)

// UserStatus struct
type UserStatus struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName returns the table name of UserStatus model
func (UserStatus) TableName() string {
	return "user_statuses"
}

// Role struct
type Role struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName returns the table name of Role model
func (Role) TableName() string {
	return "roles"
}

// User model
type User struct {
	ID           int        `gorm:"column:id;primary_key" json:"id"`
	Username     string     `gorm:"column:username;type:varchar(255)" json:"username"`
	Email        string     `gorm:"column:email;type:varchar(255)" json:"email"`
	PasswordHash string     `gorm:"column:password_hash;type:varchar(255)" json:"password_hash"`
	PasswordSalt string     `gorm:"column:password_salt;type:varchar(255)" json:"password_salt"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
	RoleID       int        `gorm:"column:role_id" json:"role_id"`
	StatusID     int        `gorm:"column:status_id" json:"status_id"`
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"role"`
	Status       UserStatus `gorm:"foreignKey:StatusID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"status"`
}

// TableName return table name of User model
func (User) TableName() string {
	return "users"
}

// UserCard struct
type UserCard struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	CardNumber string    `gorm:"column:card_number;type:varchar(255)" json:"card_number"`
	Barcode    string    `gorm:"column:barcode;type:varchar(255)" json:"barcode"`
	IssuedAt   time.Time `gorm:"column:issued_at" json:"issued_at"`
	Active     bool      `gorm:"column:active" json:"active"`
	UserID     int       `gorm:"column:user_id" json:"user_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName returns the table name of UserCard model
func (UserCard) TableName() string {
	return "user_cards"
}
