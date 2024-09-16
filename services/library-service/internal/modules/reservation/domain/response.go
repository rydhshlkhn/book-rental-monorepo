package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseReservationList model
type ResponseReservationList struct {
	Meta candishared.Meta      `json:"meta"`
	Data []ResponseReservation `json:"data"`
}

// ResponseReservation model
type ResponseReservation struct {
	ID         int    `json:"id"`
	BookItemID int    `json:"book_item_id"`
	StatusID   int    `json:"status_id"`
	UserID     int    `json:"user_id"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseReservation) Serialize(source *shareddomain.Reservation) {
	r.ID = source.ID
	r.BookItemID = source.BookItemID
	r.StatusID = source.StatusID
	r.UserID = source.UserID
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
