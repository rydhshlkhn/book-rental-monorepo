package usecase

import (
	"context"

	"monorepo/services/book-service/internal/modules/book/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) GetDetailBook(ctx context.Context, id int) (result domain.ResponseBook, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:GetDetailBook")
	defer trace.Finish()

	repoFilter := domain.FilterBook{ID: &id}
	data, err := uc.repoSQL.BookRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
