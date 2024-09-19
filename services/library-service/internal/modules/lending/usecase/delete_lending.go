package usecase

import (
	"context"

	shareddomain "monorepo/services/library-service/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *lendingUsecaseImpl) DeleteLending(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "LendingUsecase:DeleteLending")
	defer trace.Finish()

	repoFilter := shareddomain.LendingParamGet{ID: &id}
	return uc.repoSQL.LendingRepo().Delete(ctx, &repoFilter)
}
