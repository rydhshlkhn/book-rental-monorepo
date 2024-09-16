package usecase

import (
	"context"
	"errors"

	"monorepo/services/library-service/internal/modules/fine/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/fine/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_fineUsecaseImpl_UpdateFine(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		fineRepo := &mockrepo.FineRepository{}
		fineRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Fine{}, nil)
		fineRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("FineRepo").Return(fineRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := fineUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateFine(ctx, &domain.RequestFine{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		fineRepo := &mockrepo.FineRepository{}
		fineRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Fine{}, errors.New("Error"))
		fineRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("FineRepo").Return(fineRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := fineUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateFine(ctx, &domain.RequestFine{})
		assert.Error(t, err)
	})
}
