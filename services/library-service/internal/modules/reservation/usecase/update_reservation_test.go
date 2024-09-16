package usecase

import (
	"context"
	"errors"

	"monorepo/services/library-service/internal/modules/reservation/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/reservation/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_reservationUsecaseImpl_UpdateReservation(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		reservationRepo := &mockrepo.ReservationRepository{}
		reservationRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Reservation{}, nil)
		reservationRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ReservationRepo").Return(reservationRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := reservationUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateReservation(ctx, &domain.RequestReservation{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		reservationRepo := &mockrepo.ReservationRepository{}
		reservationRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Reservation{}, errors.New("Error"))
		reservationRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ReservationRepo").Return(reservationRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := reservationUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateReservation(ctx, &domain.RequestReservation{})
		assert.Error(t, err)
	})
}
