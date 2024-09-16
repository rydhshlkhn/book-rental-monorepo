package domain

import (
	"time"
)

type Fine struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	BookItemID int       `gorm:"column:book_item_id" json:"book_item_id"`
	UserID     int       `gorm:"column:user_id" json:"user_id"`
	Amount     int       `gorm:"column:amount" json:"amount"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Fine) TableName() string {
	return "fine"
}
