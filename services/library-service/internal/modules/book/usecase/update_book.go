package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) UpdateBook(ctx context.Context, data *domain.RequestBook) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:UpdateBook")
	defer trace.Finish()

	repoFilter := domain.FilterBook{ID: &data.ID}
	existing, err := uc.repoSQL.BookRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}

	existing.ISBN = data.ISBN
	existing.Title = data.Title
	existing.Subject = data.Subject
	existing.Publisher = data.Publisher
	existing.Language = data.Language
	existing.NumberOfPage = data.NumberOfPage

	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.BookRepo().SaveBook(ctx, &existing, candishared.DBUpdateSetUpdatedFields("ISBN", "Title", "Subject", "Publisher", "Language", "NumberOfPage"))
	})
	return
}
