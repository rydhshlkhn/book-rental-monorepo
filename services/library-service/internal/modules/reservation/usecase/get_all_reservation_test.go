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

func Test_reservationUsecaseImpl_GetAllReservation(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		reservationRepo := &mockrepo.ReservationRepository{}
		reservationRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Reservation{}, nil)
		reservationRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ReservationRepo").Return(reservationRepo)

		uc := reservationUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllReservation(context.Background(), &domain.FilterReservation{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		reservationRepo := &mockrepo.ReservationRepository{}
		reservationRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Reservation{}, errors.New("Error"))
		reservationRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ReservationRepo").Return(reservationRepo)

		uc := reservationUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllReservation(context.Background(), &domain.FilterReservation{})
		assert.Error(t, err)
	})
}
