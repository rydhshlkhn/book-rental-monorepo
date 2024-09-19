package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponsePaymentList model
type ResponsePaymentList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponsePayment   `json:"data"`
}

// ResponsePayment model
type ResponsePayment struct {
	ID        int `json:"id"`
	TransactionID     string `json:"transaction_id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponsePayment) Serialize(source *shareddomain.Transaction) {
	r.ID = source.ID
	r.TransactionID = source.TransactionID
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
