package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/lending/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *lendingUsecaseImpl) GetDetailLending(ctx context.Context, id int) (result domain.ResponseLending, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "LendingUsecase:GetDetailLending")
	defer trace.Finish()

	repoFilter := domain.FilterLending{ID: &id}
	data, err := uc.repoSQL.LendingRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
