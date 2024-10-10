package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/activity/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *activityUsecaseImpl) GetAllActivity(ctx context.Context, filter *domain.FilterActivity) (result domain.ResponseActivityList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ActivityUsecase:GetAllActivity")
	defer trace.Finish()

	data, err := uc.repoSQL.ActivityRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.ActivityRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseActivity, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
