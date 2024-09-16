package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/reservation/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *reservationUsecaseImpl) GetAllReservation(ctx context.Context, filter *domain.FilterReservation) (result domain.ResponseReservationList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ReservationUsecase:GetAllReservation")
	defer trace.Finish()

	data, err := uc.repoSQL.ReservationRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.ReservationRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseReservation, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
