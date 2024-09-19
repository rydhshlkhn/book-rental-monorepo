package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/payment/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_paymentUsecaseImpl_GetDetailPayment(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		paymentRepo := &mockrepo.PaymentRepository{}
		paymentRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Payment{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("PaymentRepo").Return(paymentRepo)

		uc := paymentUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailPayment(context.Background(), 1)
		assert.NoError(t, err)
	})
}
