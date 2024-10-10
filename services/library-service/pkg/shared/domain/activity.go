package domain

import (
	"time"
)

// Activity model
type Activity struct {
	ID         int    `gorm:"column:id;primary_key" json:"id"`
	Field      string    `gorm:"column:field;type:varchar(255)" json:"field"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of Activity model
func (Activity) TableName() string {
	return "activitys"
}

