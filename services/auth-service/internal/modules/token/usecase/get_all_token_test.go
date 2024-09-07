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

func Test_tokenUsecaseImpl_GetAllToken(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		tokenRepo := &mockrepo.TokenRepository{}
		tokenRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Token{}, nil)
		tokenRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("TokenRepo").Return(tokenRepo)

		uc := tokenUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllToken(context.Background(), &domain.FilterToken{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		tokenRepo := &mockrepo.TokenRepository{}
		tokenRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Token{}, errors.New("Error"))
		tokenRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("TokenRepo").Return(tokenRepo)

		uc := tokenUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllToken(context.Background(), &domain.FilterToken{})
		assert.Error(t, err)
	})
}
