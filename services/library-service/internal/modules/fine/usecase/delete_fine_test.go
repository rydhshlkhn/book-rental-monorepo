package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/fine/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_fineUsecaseImpl_DeleteFine(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		fineRepo := &mockrepo.FineRepository{}
		fineRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("FineRepo").Return(fineRepo)

		uc := fineUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteFine(context.Background(), 1)
		assert.NoError(t, err)
	})
}
