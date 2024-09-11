package usecase

import (
	"context"
	"errors"

	"monorepo/services/book-service/internal/modules/book/domain"
	mockrepo "monorepo/services/book-service/pkg/mocks/modules/book/repository"
	mocksharedrepo "monorepo/services/book-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/book-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bookUsecaseImpl_UpdateBook(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		bookRepo := &mockrepo.BookRepository{}
		bookRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Book{}, nil)
		bookRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BookRepo").Return(bookRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := bookUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateBook(ctx, &domain.RequestBook{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		bookRepo := &mockrepo.BookRepository{}
		bookRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Book{}, errors.New("Error"))
		bookRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BookRepo").Return(bookRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := bookUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateBook(ctx, &domain.RequestBook{})
		assert.Error(t, err)
	})
}
