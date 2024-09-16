package usecase

import (
	"context"
	"fmt"

	domainbook "monorepo/services/library-service/internal/modules/book/domain"
	"monorepo/services/library-service/internal/modules/lending/domain"
	domainreservation "monorepo/services/library-service/internal/modules/reservation/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *lendingUsecaseImpl) CreateLending(ctx context.Context, req *domain.RequestLending) (result domain.ResponseLending, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "LendingUsecase:CreateLending")
	defer trace.Finish()

	filter :=  domainreservation.FilterReservation{BookItemID: req.BookItemID, UserID: req.UserID}
	reservation, err := uc.repoSQL.ReservationRepo().Find(ctx, &filter)

	err = uc.repoSQL.WithTransaction(ctx, func(txCtx context.Context) error {
		// update reservation
		if reservation.ID != 0 {
			err = <-uc.repoSQL.ReservationRepo().Save(ctx, &reservation)
			if err != nil {
				return err
			}
		}

		filterBook :=  domainbook.FilterBookItem{ID: &req.BookItemID}
		bookItem, err := uc.repoSQL.BookRepo().FindItem(ctx, &filterBook)
		if err != nil {
			return err
		}
		if bookItem.StatusID != 1 {
			return fmt.Errorf("book item already loaned")
		}

		bookItem.StatusID = 3
		err = <-uc.repoSQL.BookRepo().SaveBookItem(ctx, &bookItem)
		if err != nil {
			return err
		}

		data := req.Deserialize()
		err = <- uc.repoSQL.LendingRepo().Save(ctx, &data)
		if err != nil {
			return err
		}

		result.Serialize(&data)
		
		return nil
	})
	
	return

	// if reservation.ID != 0 {
	// 	reservation.StatusID = 3
	// 	err = uc.repoSQL.ReservationRepo().Save(ctx, &reservation)
	// 	if err != nil { return }
	// }

	// filterBook :=  domainbook.FilterBookItem{ID: &req.BookItemID}
	// bookItem, err := uc.repoSQL.BookRepo().FindItem(ctx, &filterBook)
	// if err != nil { return }
	// if bookItem.StatusID != 1 {
	// 	err = fmt.Errorf("book item already loaned")
	// 	return
	// }

	// bookItem.StatusID = 3
	// err = uc.repoSQL.BookRepo().SaveBookItem(ctx, &bookItem)
	// if err != nil { return }

	// data := req.Deserialize()
	// err = uc.repoSQL.LendingRepo().Save(ctx, &data)
	// if err != nil { return }
	// result.Serialize(&data)

	// Sample using broker publisher
	// uc.deps.GetBroker(types.Kafka). // get registered broker type (sample Kafka)
	// 				GetPublisher().
	// 				PublishMessage(ctx, &candishared.PublisherArgument{
	// 		Topic:   "[topic]",
	// 		Key:     "[key]",
	// 		Message: candihelper.ToBytes("[message]"),
	// 	})
	return
}
