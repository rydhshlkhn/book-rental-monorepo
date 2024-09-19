package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/payment/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *paymentUsecaseImpl) UpdatePayment(ctx context.Context, data *domain.RequestPayment) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:UpdatePayment")
	defer trace.Finish()

	repoFilter := domain.FilterPayment{ID: &data.ID}
	existing, err := uc.repoSQL.PaymentRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.TransactionID = data.TransactionID
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.PaymentRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
