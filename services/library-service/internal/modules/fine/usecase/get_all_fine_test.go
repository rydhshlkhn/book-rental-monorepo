package usecase

import (
	"context"
	"errors"

	"monorepo/services/library-service/internal/modules/fine/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/fine/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_fineUsecaseImpl_GetAllFine(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		fineRepo := &mockrepo.FineRepository{}
		fineRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Fine{}, nil)
		fineRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("FineRepo").Return(fineRepo)

		uc := fineUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllFine(context.Background(), &domain.FilterFine{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		fineRepo := &mockrepo.FineRepository{}
		fineRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Fine{}, errors.New("Error"))
		fineRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("FineRepo").Return(fineRepo)

		uc := fineUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllFine(context.Background(), &domain.FilterFine{})
		assert.Error(t, err)
	})
}
