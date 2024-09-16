package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/fine/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *fineUsecaseImpl) GetAllFine(ctx context.Context, filter *domain.FilterFine) (result domain.ResponseFineList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineUsecase:GetAllFine")
	defer trace.Finish()

	data, err := uc.repoSQL.FineRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.FineRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseFine, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
