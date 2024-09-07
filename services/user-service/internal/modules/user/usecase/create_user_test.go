package usecase

import (
	"context"

	"monorepo/services/user-service/internal/modules/user/domain"
	mockrepo "monorepo/services/user-service/pkg/mocks/modules/user/repository"
	mocksharedrepo "monorepo/services/user-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseImpl_CreateUser(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		userRepo := &mockrepo.UserRepository{}
		userRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)

		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateUser(context.Background(), &domain.RequestUser{})
		assert.NoError(t, err)
	})
}
