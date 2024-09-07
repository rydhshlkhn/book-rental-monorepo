package domain

import (
	"time"
)

// Token model
type Token struct {
	ID           int    `gorm:"column:id;primary_key" json:"id"`
	Field        string `gorm:"column:field;type:varchar(255)" json:"field"`
	UserID       string `gorm:"column:user_id" json:"user_id"`
	DeviceID     string `gorm:"column:device_id" json:"device_id"`
	Token        string `gorm:"column:token" json:"token"`
	RefreshToken string `gorm:"column:refresh_token" json:"refresh_token"`
	IsActive     *bool  `gorm:"column:is_active" json:"is_active"`
	// Claims       map[string]interface{} `gorm:"column:claims" json:"claims"`
	ExpiredAt time.Time `gorm:"column:expired_at" json:"expired_at"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Token model
func (Token) TableName() string {
	return "tokens"
}
