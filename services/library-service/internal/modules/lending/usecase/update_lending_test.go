package usecase

import (
	"context"
	"errors"

	"monorepo/services/library-service/internal/modules/lending/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/lending/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_lendingUsecaseImpl_UpdateLending(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		lendingRepo := &mockrepo.LendingRepository{}
		lendingRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Lending{}, nil)
		lendingRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("LendingRepo").Return(lendingRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := lendingUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateLending(ctx, &domain.RequestLending{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		lendingRepo := &mockrepo.LendingRepository{}
		lendingRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Lending{}, errors.New("Error"))
		lendingRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("LendingRepo").Return(lendingRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := lendingUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateLending(ctx, &domain.RequestLending{})
		assert.Error(t, err)
	})
}
