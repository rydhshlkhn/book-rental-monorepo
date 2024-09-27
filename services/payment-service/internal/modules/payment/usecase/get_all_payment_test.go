package usecase

import (
	"context"
	"errors"

	"monorepo/services/payment-service/internal/modules/payment/domain"
	mockrepo "monorepo/services/payment-service/pkg/mocks/modules/payment/repository"
	mocksharedrepo "monorepo/services/payment-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/payment-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_paymentUsecaseImpl_GetAllPayment(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		paymentRepo := &mockrepo.PaymentRepository{}
		paymentRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Payment{}, nil)
		paymentRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("PaymentRepo").Return(paymentRepo)

		uc := paymentUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllPayment(context.Background(), &domain.FilterPayment{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		paymentRepo := &mockrepo.PaymentRepository{}
		paymentRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Payment{}, errors.New("Error"))
		paymentRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("PaymentRepo").Return(paymentRepo)

		uc := paymentUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllPayment(context.Background(), &domain.FilterPayment{})
		assert.Error(t, err)
	})
}
