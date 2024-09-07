package usecase

import (
	"context"
	"errors"

	"monorepo/services/user-service/internal/modules/user/domain"
	mockrepo "monorepo/services/user-service/pkg/mocks/modules/user/repository"
	mocksharedrepo "monorepo/services/user-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/user-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseImpl_UpdateUser(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		userRepo := &mockrepo.UserRepository{}
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, nil)
		userRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateUser(ctx, &domain.RequestUser{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		userRepo := &mockrepo.UserRepository{}
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, errors.New("Error"))
		userRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateUser(ctx, &domain.RequestUser{})
		assert.Error(t, err)
	})
}
