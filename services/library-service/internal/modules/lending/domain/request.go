package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"time"
)

// RequestLending model
type RequestLending struct {
	ID         int       `json:"id"`
	BookItemID int       `json:"book_item_id"`
	UserID     int       `json:"user_id"`
	DueDate    time.Time `json:"due_date"`
	ReturnDate time.Time `json:"return_date"`
}

// Deserialize to db model
func (r *RequestLending) Deserialize() (res shareddomain.Lending) {
	res.ID = r.ID
	res.BookItemID = r.BookItemID
	res.UserID = r.UserID
	res.DueDate = r.DueDate
	res.ReturnDate = r.ReturnDate
	return
}
