package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/fine/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_fineUsecaseImpl_GetDetailFine(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		fineRepo := &mockrepo.FineRepository{}
		fineRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Fine{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("FineRepo").Return(fineRepo)

		uc := fineUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailFine(context.Background(), 1)
		assert.NoError(t, err)
	})
}
