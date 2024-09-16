package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/reservation/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *reservationUsecaseImpl) UpdateReservation(ctx context.Context, data *domain.RequestReservation) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ReservationUsecase:UpdateReservation")
	defer trace.Finish()

	repoFilter := domain.FilterReservation{ID: &data.ID}
	existing, err := uc.repoSQL.ReservationRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.BookItemID = data.BookItemID
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return <-uc.repoSQL.ReservationRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
