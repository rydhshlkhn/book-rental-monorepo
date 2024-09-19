package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/payment/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *paymentUsecaseImpl) GetAllPayment(ctx context.Context, filter *domain.FilterPayment) (result domain.ResponsePaymentList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:GetAllPayment")
	defer trace.Finish()

	data, err := uc.repoSQL.PaymentRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.PaymentRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponsePayment, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
