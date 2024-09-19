package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/payment/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_paymentUsecaseImpl_DeletePayment(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		paymentRepo := &mockrepo.PaymentRepository{}
		paymentRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("PaymentRepo").Return(paymentRepo)

		uc := paymentUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeletePayment(context.Background(), 1)
		assert.NoError(t, err)
	})
}
