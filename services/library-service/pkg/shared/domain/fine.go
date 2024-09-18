package domain

import (
	"time"
)

type Fine struct {
	ID            int       `gorm:"column:id;primary_key" json:"id"`
	LendingID     int       `gorm:"column:lending_id" json:"lending_id"`
	Amount        int       `gorm:"column:amount" json:"amount"`
	SnanpURL      string    `gorm:"column:snap_url" json:"snap_url"`
	PaymentStatus string    `gorm:"column:payment_status" json:"payment_status"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Fine) TableName() string {
	return "fine"
}
