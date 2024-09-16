package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/lending/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_lendingUsecaseImpl_GetDetailLending(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		lendingRepo := &mockrepo.LendingRepository{}
		lendingRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Lending{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("LendingRepo").Return(lendingRepo)

		uc := lendingUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailLending(context.Background(), 1)
		assert.NoError(t, err)
	})
}
