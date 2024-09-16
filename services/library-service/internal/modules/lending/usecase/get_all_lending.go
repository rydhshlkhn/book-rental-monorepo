package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/lending/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *lendingUsecaseImpl) GetAllLending(ctx context.Context, filter *domain.FilterLending) (result domain.ResponseLendingList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "LendingUsecase:GetAllLending")
	defer trace.Finish()

	data, err := uc.repoSQL.LendingRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.LendingRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseLending, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
