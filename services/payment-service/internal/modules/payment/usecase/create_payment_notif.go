package usecase

import (
	"context"
	"fmt"

	"monorepo/services/payment-service/internal/modules/payment/domain"
	"monorepo/services/payment-service/pkg/helper"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *paymentUsecaseImpl) CreatePaymentNotif(ctx context.Context, req *domain.RequestTransaction) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:CreatePaymentNotif")
	defer trace.Finish()

	data, err := req.Deserialize()
	if err != nil {
		return
	}

	fmt.Println(data)

	if isValidSignature := helper.VerifyingSignatureKey(data.SignatureKey, data.OrderID, data.StatusCode, data.GrossAmount); !isValidSignature {
		return fmt.Errorf("invalid signature key")
	}

	err = uc.repoSQL.PaymentRepo().SaveNotif(ctx, &data)
	if err != nil {
		return
	}

	// Sample using broker publisher
	uc.publisher.PublishMessage(ctx, &candishared.PublisherArgument{
		Topic:   "lending",
		Key:     data.OrderID,
		Message: candihelper.ToBytes(data),
	})
	return
}
