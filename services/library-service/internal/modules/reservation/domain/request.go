package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
)

// RequestReservation model
type RequestReservation struct {
	ID         int `json:"id"`
	BookItemID int `json:"book_item_id"`
	StatusID   int `json:"status_id"`
	UserID     int `json:"user_id"`
}

// Deserialize to db model
func (r *RequestReservation) Deserialize() (res shareddomain.Reservation) {
	res.BookItemID = r.BookItemID
	res.StatusID = r.StatusID
	res.UserID = r.UserID
	return
}
