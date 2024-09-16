package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/reservation/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *reservationUsecaseImpl) GetDetailReservation(ctx context.Context, id int) (result domain.ResponseReservation, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ReservationUsecase:GetDetailReservation")
	defer trace.Finish()

	repoFilter := domain.FilterReservation{ID: &id}
	data, err := uc.repoSQL.ReservationRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
