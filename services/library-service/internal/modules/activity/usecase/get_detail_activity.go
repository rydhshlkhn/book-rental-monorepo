package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/activity/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *activityUsecaseImpl) GetDetailActivity(ctx context.Context, id int) (result domain.ResponseActivity, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ActivityUsecase:GetDetailActivity")
	defer trace.Finish()

	repoFilter := domain.FilterActivity{ID: &id}
	data, err := uc.repoSQL.ActivityRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
