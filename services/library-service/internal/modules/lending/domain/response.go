package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseLendingList model
type ResponseLendingList struct {
	Meta candishared.Meta  `json:"meta"`
	Data []ResponseLending `json:"data"`
}

// ResponseLending model
type ResponseLending struct {
	ID         int    `json:"id"`
	BookItemID int    `json:"book_item_id"`
	DueDate    string `json:"due_date"`
	ReturnDate string `json:"return_date"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseLending) Serialize(source *shareddomain.Lending) {
	r.ID = source.ID
	r.BookItemID = source.BookItemID
	r.DueDate = source.DueDate.Format(time.RFC3339)
	r.ReturnDate = source.ReturnDate.Format(time.RFC3339)
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
