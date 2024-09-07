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

func Test_userUsecaseImpl_GetAllUser(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		userRepo := &mockrepo.UserRepository{}
		userRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.User{}, nil)
		userRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)

		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllUser(context.Background(), &domain.FilterUser{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		userRepo := &mockrepo.UserRepository{}
		userRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.User{}, errors.New("Error"))
		userRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)

		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllUser(context.Background(), &domain.FilterUser{})
		assert.Error(t, err)
	})
}
