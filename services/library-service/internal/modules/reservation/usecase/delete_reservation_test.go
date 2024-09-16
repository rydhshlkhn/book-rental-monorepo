package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/reservation/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_reservationUsecaseImpl_DeleteReservation(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		reservationRepo := &mockrepo.ReservationRepository{}
		reservationRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ReservationRepo").Return(reservationRepo)

		uc := reservationUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteReservation(context.Background(), 1)
		assert.NoError(t, err)
	})
}
