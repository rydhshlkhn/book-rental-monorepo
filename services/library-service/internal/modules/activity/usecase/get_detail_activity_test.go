package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/activity/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_activityUsecaseImpl_GetDetailActivity(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		activityRepo := &mockrepo.ActivityRepository{}
		activityRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Activity{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ActivityRepo").Return(activityRepo)

		uc := activityUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailActivity(context.Background(), 1)
		assert.NoError(t, err)
	})
}
