package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
)

// RequestFine model
type RequestFine struct {
	ID         int `json:"id"`
	LendingID int `json:"lending_id"`
	PaymentStatus string `json:"payment_status"`
	Amount     int `json:"amount"`
}

// Deserialize to db model
func (r *RequestFine) Deserialize() (res shareddomain.Fine) {
	res.ID = r.ID
	res.LendingID = r.LendingID
	res.Amount = r.Amount
	res.PaymentStatus = r.PaymentStatus
	return
}
