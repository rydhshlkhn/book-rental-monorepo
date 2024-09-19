package domain

import (
	"monorepo/services/library-service/pkg/helper"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
)

// RequestPayment model
type RequestPayment struct {
	ID            int    `json:"id"`
	TransactionID string `json:"transaction_id"`
}

// Deserialize to db model
func (r *RequestPayment) Deserialize() (res shareddomain.Transaction) {
	res.TransactionID = r.TransactionID
	return
}

// Transaction model
type RequestTransaction struct {
	ID                int             `json:"id"`
	VANumbers         []VANumber      `json:"va_numbers"`
	TransactionTime   string          `json:"transaction_time"`
	TransactionStatus string          `json:"transaction_status"`
	TransactionID     string          `json:"transaction_id"`
	StatusMessage     string          `json:"status_message"`
	StatusCode        string          `json:"status_code"`
	SignatureKey      string          `json:"signature_key"`
	SettlementTime    string          `json:"settlement_time"`
	PaymentType       string          `json:"payment_type"`
	PaymentAmounts    []PaymentAmount `json:"payment_amounts"`
	OrderID           string          `json:"order_id"`
	MerchantID        string          `json:"merchant_id"`
	GrossAmount       string          `json:"gross_amount"`
	FraudStatus       string          `json:"fraud_status"`
	Currency          string          `json:"currency"`
}

// VANumber model
type VANumber struct {
	VANumber string `json:"va_number"`
	Bank     string `json:"bank"`
}

// PaymentAmount model
type PaymentAmount struct {
	PaidAt string `json:"paid_at"`
	Amount string `json:"amount"`
}

// Deserialize to db model
func (t *RequestTransaction) Deserialize() (res shareddomain.Transaction, err error) {
	res.VANumbers = make([]shareddomain.VANumber, len(t.VANumbers))
	for i, va := range t.VANumbers {
		res.VANumbers[i] = shareddomain.VANumber{
			VANumber: va.VANumber,
			Bank:     va.Bank,
		}
	}

	transatcionTime, err := helper.ParseDate(t.TransactionTime)
	if err != nil {
		return
	}
	settlementTime, err := helper.ParseDate(t.SettlementTime)
	if err != nil {
		return
	}

	res.TransactionTime = transatcionTime
	res.TransactionStatus = t.TransactionStatus
	res.TransactionID = t.TransactionID
	res.StatusMessage = t.StatusMessage
	res.StatusCode = t.StatusCode
	res.SignatureKey = t.SignatureKey
	res.SettlementTime = settlementTime
	res.PaymentType = t.PaymentType

	res.PaymentAmounts = make([]shareddomain.PaymentAmount, len(t.PaymentAmounts))
	for i, p := range t.PaymentAmounts {
		paidAt, err := helper.ParseDate(p.PaidAt)
		if err != nil {
			return res, err
		}

		res.PaymentAmounts[i] = shareddomain.PaymentAmount{
			PaidAt: paidAt,
			Amount: p.Amount,
		}
	}

	res.OrderID = t.OrderID
	res.MerchantID = t.MerchantID
	res.GrossAmount = t.GrossAmount
	res.FraudStatus = t.FraudStatus
	res.Currency = t.Currency
	return
}
