package usecase

import (
	"context"

	mockrepo "monorepo/services/auth-service/pkg/mocks/modules/token/repository"
	mocksharedrepo "monorepo/services/auth-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_tokenUsecaseImpl_DeleteToken(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		tokenRepo := &mockrepo.TokenRepository{}
		tokenRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("TokenRepo").Return(tokenRepo)

		uc := tokenUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteToken(context.Background(), 1)
		assert.NoError(t, err)
	})
}
