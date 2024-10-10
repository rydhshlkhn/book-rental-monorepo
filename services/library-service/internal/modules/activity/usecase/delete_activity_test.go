package usecase

import (
	"context"

	mockrepo "monorepo/services/library-service/pkg/mocks/modules/activity/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_activityUsecaseImpl_DeleteActivity(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		activityRepo := &mockrepo.ActivityRepository{}
		activityRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ActivityRepo").Return(activityRepo)

		uc := activityUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteActivity(context.Background(), 1)
		assert.NoError(t, err)
	})
}
