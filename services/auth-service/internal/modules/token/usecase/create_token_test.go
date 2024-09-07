package usecase

import (
	"context"

	"monorepo/services/auth-service/internal/modules/token/domain"
	mockrepo "monorepo/services/auth-service/pkg/mocks/modules/token/repository"
	mocksharedrepo "monorepo/services/auth-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_tokenUsecaseImpl_CreateToken(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		tokenRepo := &mockrepo.TokenRepository{}
		tokenRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("TokenRepo").Return(tokenRepo)

		uc := tokenUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateToken(context.Background(), &domain.RequestToken{})
		assert.NoError(t, err)
	})
}
