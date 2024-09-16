package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) UpdateBookItem(ctx context.Context, data *domain.RequestBookItem) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:UpdateBookItem")
	defer trace.Finish()

	repoFilter := domain.FilterBookItem{ID: &data.ID}
	existing, err := uc.repoSQL.BookRepo().FindItem(ctx, &repoFilter)
	if err != nil {
		return err
	}

	existing.Barcode = data.Barcode
	existing.IsReferenceOnly = data.IsReferenceOnly
	existing.Borrowed = data.Borrowed
	existing.DueDate = data.DueDate
	existing.FormatID = data.FormatID
	existing.StatusID = data.StatusID
	existing.DateOfPurchase = data.DateOfPurchase
	existing.PublicationDate = data.PublicationDate
	
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return <-uc.repoSQL.BookRepo().SaveBookItem(ctx, &existing, candishared.DBUpdateSetUpdatedFields(
			"Barcode", "IsReferenceOnly", "Borrowed", "DueDate", 
			"FormatID", "StatusID", "DateOfPurchase", "PublicationDate",
		))
	})
	return
}
