package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/payment/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *paymentUsecaseImpl) GetDetailPayment(ctx context.Context, id int) (result domain.ResponsePayment, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:GetDetailPayment")
	defer trace.Finish()

	repoFilter := domain.FilterPayment{ID: &id}
	data, err := uc.repoSQL.PaymentRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
