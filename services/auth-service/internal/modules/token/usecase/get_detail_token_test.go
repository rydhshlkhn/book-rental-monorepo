package usecase

import (
	"context"

	mockrepo "monorepo/services/auth-service/pkg/mocks/modules/token/repository"
	mocksharedrepo "monorepo/services/auth-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/auth-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_tokenUsecaseImpl_GetDetailToken(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		tokenRepo := &mockrepo.TokenRepository{}
		tokenRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Token{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("TokenRepo").Return(tokenRepo)

		uc := tokenUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailToken(context.Background(), 1)
		assert.NoError(t, err)
	})
}
