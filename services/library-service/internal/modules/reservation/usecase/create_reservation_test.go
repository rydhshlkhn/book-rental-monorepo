package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/reservation/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/reservation/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_reservationUsecaseImpl_CreateReservation(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		reservationRepo := &mockrepo.ReservationRepository{}
		reservationRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ReservationRepo").Return(reservationRepo)

		uc := reservationUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateReservation(context.Background(), &domain.RequestReservation{})
		assert.NoError(t, err)
	})
}
