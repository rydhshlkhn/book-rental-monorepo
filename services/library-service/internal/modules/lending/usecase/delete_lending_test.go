package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/lending/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_lendingUsecaseImpl_DeleteLending(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		lendingRepo := &mockrepo.LendingRepository{}
		lendingRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("LendingRepo").Return(lendingRepo)

		uc := lendingUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteLending(context.Background(), 1)
		assert.NoError(t, err)
	})
}
