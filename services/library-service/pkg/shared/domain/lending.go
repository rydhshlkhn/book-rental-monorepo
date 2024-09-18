package domain

import (
	"time"
)

type Lending struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	BookItemID int       `gorm:"column:book_item_id" json:"book_item_id"`
	UserID     int       `gorm:"column:user_id" json:"user_id"`
	DueDate    time.Time `gorm:"column:due_date" json:"due_date"`
	ReturnDate *time.Time `gorm:"column:return_date" json:"return_date"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Lending) TableName() string {
	return "lending"
}
