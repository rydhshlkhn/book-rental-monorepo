package usecase

import (
	"context"
	"errors"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/lending/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_lendingUsecaseImpl_GetAllLending(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		lendingRepo := &mockrepo.LendingRepository{}
		lendingRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Lending{}, nil)
		lendingRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("LendingRepo").Return(lendingRepo)

		uc := lendingUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllLending(context.Background(), &shareddomain.LendingParamGet{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		lendingRepo := &mockrepo.LendingRepository{}
		lendingRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Lending{}, errors.New("Error"))
		lendingRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("LendingRepo").Return(lendingRepo)

		uc := lendingUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllLending(context.Background(), &shareddomain.LendingParamGet{})
		assert.Error(t, err)
	})
}
