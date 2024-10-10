package usecase

import (
	"context"
	"errors"

	"monorepo/services/library-service/internal/modules/activity/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/activity/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_activityUsecaseImpl_GetAllActivity(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		activityRepo := &mockrepo.ActivityRepository{}
		activityRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Activity{}, nil)
		activityRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ActivityRepo").Return(activityRepo)

		uc := activityUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllActivity(context.Background(), &domain.FilterActivity{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		activityRepo := &mockrepo.ActivityRepository{}
		activityRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Activity{}, errors.New("Error"))
		activityRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ActivityRepo").Return(activityRepo)

		uc := activityUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllActivity(context.Background(), &domain.FilterActivity{})
		assert.Error(t, err)
	})
}
