package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/fine/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *fineUsecaseImpl) UpdateFine(ctx context.Context, data *domain.RequestFine) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineUsecase:UpdateFine")
	defer trace.Finish()

	repoFilter := domain.FilterFine{ID: &data.ID}
	existing, err := uc.repoSQL.FineRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.BookItemID = data.BookItemID
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.FineRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
