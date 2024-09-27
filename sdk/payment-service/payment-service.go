package paymentservice

import "context"

// Paymentservice client abstract interface
type Paymentservice interface {
	// Add service client method
	Pay(ctx context.Context, req RequestPayment) (res *Response, code int, err error)
}
