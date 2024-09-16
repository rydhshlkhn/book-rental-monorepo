package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) GetAllBook(ctx context.Context, filter *domain.FilterBook) (result domain.ResponseBookList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:GetAllBook")
	defer trace.Finish()

	data, err := uc.repoSQL.BookRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.BookRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseBook, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
