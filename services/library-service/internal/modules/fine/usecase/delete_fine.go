package usecase

import (
	"context"
	
	"monorepo/services/library-service/internal/modules/fine/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *fineUsecaseImpl) DeleteFine(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineUsecase:DeleteFine")
	defer trace.Finish()

	repoFilter := domain.FilterFine{ID: &id}
	return uc.repoSQL.FineRepo().Delete(ctx, &repoFilter)
}
