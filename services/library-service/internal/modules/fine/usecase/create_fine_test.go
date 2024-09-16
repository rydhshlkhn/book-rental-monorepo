package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/fine/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/fine/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_fineUsecaseImpl_CreateFine(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		fineRepo := &mockrepo.FineRepository{}
		fineRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("FineRepo").Return(fineRepo)

		uc := fineUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateFine(context.Background(), &domain.RequestFine{})
		assert.NoError(t, err)
	})
}
