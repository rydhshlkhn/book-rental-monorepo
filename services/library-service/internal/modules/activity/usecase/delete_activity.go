package usecase

import (
	"context"
	
	"monorepo/services/library-service/internal/modules/activity/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *activityUsecaseImpl) DeleteActivity(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ActivityUsecase:DeleteActivity")
	defer trace.Finish()

	repoFilter := domain.FilterActivity{ID: &id}
	return uc.repoSQL.ActivityRepo().Delete(ctx, &repoFilter)
}
