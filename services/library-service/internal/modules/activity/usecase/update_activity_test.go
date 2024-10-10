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

func Test_activityUsecaseImpl_UpdateActivity(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		activityRepo := &mockrepo.ActivityRepository{}
		activityRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Activity{}, nil)
		activityRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ActivityRepo").Return(activityRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := activityUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateActivity(ctx, &domain.RequestActivity{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		activityRepo := &mockrepo.ActivityRepository{}
		activityRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Activity{}, errors.New("Error"))
		activityRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ActivityRepo").Return(activityRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := activityUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateActivity(ctx, &domain.RequestActivity{})
		assert.Error(t, err)
	})
}
