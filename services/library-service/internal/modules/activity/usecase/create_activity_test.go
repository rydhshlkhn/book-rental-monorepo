package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/activity/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/activity/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_activityUsecaseImpl_CreateActivity(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		activityRepo := &mockrepo.ActivityRepository{}
		activityRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ActivityRepo").Return(activityRepo)

		uc := activityUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateActivity(context.Background(), &domain.RequestActivity{})
		assert.NoError(t, err)
	})
}
