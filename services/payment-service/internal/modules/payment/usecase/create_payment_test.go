package usecase

import (
	"context"

	"monorepo/services/payment-service/internal/modules/payment/domain"
	mockrepo "monorepo/services/payment-service/pkg/mocks/modules/payment/repository"
	mocksharedrepo "monorepo/services/payment-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_paymentUsecaseImpl_CreatePayment(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		paymentRepo := &mockrepo.PaymentRepository{}
		paymentRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("PaymentRepo").Return(paymentRepo)

		uc := paymentUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreatePayment(context.Background(), &domain.RequestPayment{})
		assert.NoError(t, err)
	})
}
