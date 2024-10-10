package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/activity/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *activityUsecaseImpl) UpdateActivity(ctx context.Context, data *domain.RequestActivity) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ActivityUsecase:UpdateActivity")
	defer trace.Finish()

	repoFilter := domain.FilterActivity{ID: &data.ID}
	existing, err := uc.repoSQL.ActivityRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Field = data.Field
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.ActivityRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
