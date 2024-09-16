package usecase

import (
	"context"
	
	"monorepo/services/library-service/internal/modules/lending/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *lendingUsecaseImpl) DeleteLending(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "LendingUsecase:DeleteLending")
	defer trace.Finish()

	repoFilter := domain.FilterLending{ID: &id}
	return uc.repoSQL.LendingRepo().Delete(ctx, &repoFilter)
}
