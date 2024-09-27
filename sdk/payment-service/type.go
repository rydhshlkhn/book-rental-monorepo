package paymentservice

type RequestPayment struct {
	OrderID string `json:"order_id"`
	GrossAmount string `json:"gross_amount"`
}

type ResponsePayment struct {
	PaymentURL string `json:"payment_url"`
}

type Response struct {
	Success bool           `json:"success"`
	Data    *ResponsePayment `json:"data"`
	Message string         `json:"message"`
	Code    int            `json:"code"`
}