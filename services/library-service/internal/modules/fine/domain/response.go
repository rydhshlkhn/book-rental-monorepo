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
	ID        int    `json:"id"`
	LendingID int    `json:"lending_id"`
	Amount    int    `json:"amount"`
	SnapURL    string    `json:"snap_url"`
	PaymentStatus    string    `json:"payment_status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseFine) Serialize(source *shareddomain.Fine) {
	r.ID = source.ID
	r.LendingID = source.LendingID
	r.Amount = source.Amount
	r.SnapURL = source.SnanpURL
	r.PaymentStatus = source.PaymentStatus
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
