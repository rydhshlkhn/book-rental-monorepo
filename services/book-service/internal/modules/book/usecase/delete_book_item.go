package usecase

import (
	"context"
	
	"monorepo/services/book-service/internal/modules/book/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) DeleteBookItem(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:DeleteBook")
	defer trace.Finish()

	repoFilter := domain.FilterBookItem{ID: &id}
	return uc.repoSQL.BookRepo().DeleteItem(ctx, &repoFilter)
}
