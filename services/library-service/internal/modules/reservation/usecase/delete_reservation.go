package usecase

import (
	"context"
	
	"monorepo/services/library-service/internal/modules/reservation/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *reservationUsecaseImpl) DeleteReservation(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ReservationUsecase:DeleteReservation")
	defer trace.Finish()

	repoFilter := domain.FilterReservation{ID: &id}
	return uc.repoSQL.ReservationRepo().Delete(ctx, &repoFilter)
}
