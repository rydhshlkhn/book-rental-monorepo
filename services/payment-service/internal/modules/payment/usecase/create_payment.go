package usecase

import (
	"context"

	"monorepo/services/payment-service/internal/modules/payment/domain"
	"monorepo/services/payment-service/pkg/shared"

	"github.com/golangid/candi/tracer"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (uc *paymentUsecaseImpl) CreatePayment(ctx context.Context, req *domain.RequestPayment) (result *domain.ResponsePayment, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:CreatePayment")
	defer trace.Finish()

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(shared.GetEnv().MidtransServerKey, midtrans.Sandbox)

	// 2. Initiate Snap request
	midtransREq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: req.GrossAmount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, midErr := s.CreateTransaction(midtransREq)
	if midErr != nil {
		return nil, midErr.RawError
	}

	result = new(domain.ResponsePayment)
	result.PaymentUrl = snapResp.RedirectURL

	return
}
