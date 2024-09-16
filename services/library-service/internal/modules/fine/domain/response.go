package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseFineList model
type ResponseFineList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseFine   `json:"data"`
}

// ResponseFine model
type ResponseFine struct {
	ID         int    `json:"id"`
	BookItemID int    `json:"book_item_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseFine) Serialize(source *shareddomain.Fine) {
	r.ID = source.ID
	r.BookItemID = source.BookItemID
	r.UserID = source.UserID
	r.Amount = source.Amount
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
