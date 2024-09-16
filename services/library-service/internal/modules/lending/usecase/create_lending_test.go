package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/lending/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/lending/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_lendingUsecaseImpl_CreateLending(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		lendingRepo := &mockrepo.LendingRepository{}
		lendingRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("LendingRepo").Return(lendingRepo)

		uc := lendingUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateLending(context.Background(), &domain.RequestLending{})
		assert.NoError(t, err)
	})
}
