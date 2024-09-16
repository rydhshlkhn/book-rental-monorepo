package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/fine/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *fineUsecaseImpl) GetDetailFine(ctx context.Context, id int) (result domain.ResponseFine, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineUsecase:GetDetailFine")
	defer trace.Finish()

	repoFilter := domain.FilterFine{ID: &id}
	data, err := uc.repoSQL.FineRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
