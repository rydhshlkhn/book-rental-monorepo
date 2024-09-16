package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/reservation/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_reservationUsecaseImpl_GetDetailReservation(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		reservationRepo := &mockrepo.ReservationRepository{}
		reservationRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Reservation{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ReservationRepo").Return(reservationRepo)

		uc := reservationUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailReservation(context.Background(), 1)
		assert.NoError(t, err)
	})
}
