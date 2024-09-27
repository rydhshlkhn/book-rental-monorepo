package domain

import (
	shareddomain "monorepo/services/payment-service/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
)

// ResponsePaymentList model
type ResponsePaymentList struct {
	Meta candishared.Meta  `json:"meta"`
	Data []ResponsePayment `json:"data"`
}

// ResponsePayment model
type ResponsePayment struct {
	PaymentUrl string `json:"payment_url"`
}

// Serialize from db model
func (r *ResponsePayment) Serialize(source *shareddomain.Transaction) {
}
