package usecase

import (
	"context"
	"errors"

	"monorepo/services/auth-service/internal/modules/token/domain"
	mockrepo "monorepo/services/auth-service/pkg/mocks/modules/token/repository"
	mocksharedrepo "monorepo/services/auth-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/auth-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_tokenUsecaseImpl_UpdateToken(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		tokenRepo := &mockrepo.TokenRepository{}
		tokenRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Token{}, nil)
		tokenRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("TokenRepo").Return(tokenRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := tokenUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateToken(ctx, &domain.RequestToken{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		tokenRepo := &mockrepo.TokenRepository{}
		tokenRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Token{}, errors.New("Error"))
		tokenRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("TokenRepo").Return(tokenRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := tokenUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateToken(ctx, &domain.RequestToken{})
		assert.Error(t, err)
	})
}
