package usecase

import (
	"context"
	
	"monorepo/services/library-service/internal/modules/payment/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *paymentUsecaseImpl) DeletePayment(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:DeletePayment")
	defer trace.Finish()

	repoFilter := domain.FilterPayment{ID: &id}
	return uc.repoSQL.PaymentRepo().Delete(ctx, &repoFilter)
}
