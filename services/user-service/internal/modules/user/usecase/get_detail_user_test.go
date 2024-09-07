package usecase

import (
	"context"

	mockrepo "monorepo/services/user-service/pkg/mocks/modules/user/repository"
	mocksharedrepo "monorepo/services/user-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/user-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseImpl_GetDetailUser(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		userRepo := &mockrepo.UserRepository{}
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)

		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailUser(context.Background(), 1)
		assert.NoError(t, err)
	})
}
