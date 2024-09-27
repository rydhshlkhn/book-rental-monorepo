package domain

import (
	"time"
)

// Transaction model
type Transaction struct {
	ID                int       `gorm:"column:id;primary_key" json:"id"`
	TransactionTime   time.Time `gorm:"column:transaction_time" json:"transaction_time"`
	TransactionStatus string    `gorm:"column:transaction_status;type:varchar(50)" json:"transaction_status"`
	TransactionID     string    `gorm:"column:transaction_id;type:varchar(100)" json:"transaction_id"`
	StatusMessage     string    `gorm:"column:status_message;type:varchar(255)" json:"status_message"`
	StatusCode        string    `gorm:"column:status_code;type:varchar(10)" json:"status_code"`
	SignatureKey      string    `gorm:"column:signature_key;type:text" json:"signature_key"`
	SettlementTime    time.Time `gorm:"column:settlement_time" json:"settlement_time"`
	PaymentType       string    `gorm:"column:payment_type;type:varchar(50)" json:"payment_type"`
	OrderID           string    `gorm:"column:order_id;type:varchar(100)" json:"order_id"`
	MerchantID        string    `gorm:"column:merchant_id;type:varchar(100)" json:"merchant_id"`
	GrossAmount       string    `gorm:"column:gross_amount;type:decimal(20,2)" json:"gross_amount"`
	FraudStatus       string    `gorm:"column:fraud_status;type:varchar(50)" json:"fraud_status"`
	Currency          string    `gorm:"column:currency;type:varchar(10)" json:"currency"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relatoin One-to-One on VaNumber and PaymentAmount
	VANumbers      []VANumber      `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"va_numbers"`
	PaymentAmounts []PaymentAmount `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"payment_amounts"`
}

// VaNumber model
type VANumber struct {
	ID            int       `gorm:"column:id;primary_key" json:"id"`
	TransactionID int       `gorm:"column:transaction_id" json:"transaction_id"`
	VANumber      string    `gorm:"column:va_number;type:varchar(50)" json:"va_number"`
	Bank          string    `gorm:"column:bank;type:varchar(50)" json:"bank"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// PaymentAmount model
type PaymentAmount struct {
	ID            int       `gorm:"column:id;primary_key" json:"id"`
	TransactionID int       `gorm:"column:transaction_id" json:"transaction_id"`
	PaidAt        time.Time `gorm:"column:paid_at" json:"paid_at"`
	Amount        string    `gorm:"column:amount;type:decimal(20,2)" json:"amount"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName for Transaction
func (Transaction) TableName() string {
	return "transaction"
}

// TableName for VaNumber
func (VANumber) TableName() string {
	return "va_number"
}

// TableName for PaymentAmount
func (PaymentAmount) TableName() string {
	return "payment_amount"
}
