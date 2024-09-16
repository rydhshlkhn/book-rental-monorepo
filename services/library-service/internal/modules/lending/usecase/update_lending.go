package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/lending/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *lendingUsecaseImpl) UpdateLending(ctx context.Context, data *domain.RequestLending) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "LendingUsecase:UpdateLending")
	defer trace.Finish()

	repoFilter := domain.FilterLending{ID: &data.ID}
	existing, err := uc.repoSQL.LendingRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.BookItemID = data.BookItemID
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return <- uc.repoSQL.LendingRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
