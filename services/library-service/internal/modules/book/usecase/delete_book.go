package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) DeleteBook(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:DeleteBook")
	defer trace.Finish()

	repoFilter := domain.FilterBook{ID: &id}
	return uc.repoSQL.BookRepo().Delete(ctx, &repoFilter)
}
