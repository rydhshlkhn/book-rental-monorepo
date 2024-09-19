package usecase

import (
	"context"
	"errors"

	"monorepo/services/library-service/internal/modules/payment/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/payment/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_paymentUsecaseImpl_UpdatePayment(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		paymentRepo := &mockrepo.PaymentRepository{}
		paymentRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Payment{}, nil)
		paymentRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("PaymentRepo").Return(paymentRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := paymentUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdatePayment(ctx, &domain.RequestPayment{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		paymentRepo := &mockrepo.PaymentRepository{}
		paymentRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Payment{}, errors.New("Error"))
		paymentRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("PaymentRepo").Return(paymentRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := paymentUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdatePayment(ctx, &domain.RequestPayment{})
		assert.Error(t, err)
	})
}
