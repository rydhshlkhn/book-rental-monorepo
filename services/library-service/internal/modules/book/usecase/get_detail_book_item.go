package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) GetDetailBookItem(ctx context.Context, id int) (result domain.ResponseBookItem, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:GetDetailBookItem")
	defer trace.Finish()

	repoFilter := domain.FilterBookItem{ID: &id}
	data, err := uc.repoSQL.BookRepo().FindItem(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
