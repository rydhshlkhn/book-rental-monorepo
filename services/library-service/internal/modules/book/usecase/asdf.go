package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) Asdf(ctx context.Context, req *domain.RequestAsdf) (resp domain.ResponseAsdf, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:Asdf")
	defer trace.Finish()

	return
}
