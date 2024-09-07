package usecase

import (
	"context"

	mockrepo "monorepo/services/user-service/pkg/mocks/modules/user/repository"
	mocksharedrepo "monorepo/services/user-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseImpl_DeleteUser(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		userRepo := &mockrepo.UserRepository{}
		userRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)

		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteUser(context.Background(), 1)
		assert.NoError(t, err)
	})
}
