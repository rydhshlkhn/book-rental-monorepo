package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
)

// RequestFine model
type RequestFine struct {
	ID         int `json:"id"`
	BookItemID int `json:"book_item_id"`
	UserID     int `json:"user_id"`
	Amount     int `json:"amount"`
}

// Deserialize to db model
func (r *RequestFine) Deserialize() (res shareddomain.Fine) {
	res.ID = r.ID
	res.BookItemID = r.BookItemID
	res.UserID = r.UserID
	res.Amount = r.Amount
	return
}
