package domain

import (
	"time"
)

type ReservationStatus struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(100)" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ReservationStatus) TableName() string {
	return "reservation_status"
}

type Reservation struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	BookItemID int       `gorm:"column:book_item_id" json:"book_item_id"`
	StatusID   int       `gorm:"column:status_id" json:"status_id"`
	UserID     int       `gorm:"column:user_id" json:"user_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Reservation) TableName() string {
	return "reservation"
}
