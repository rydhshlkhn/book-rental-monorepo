package usecase

import (
	"context"
	"errors"

	"monorepo/services/library-service/internal/modules/book/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/book/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bookUsecaseImpl_GetAllBook(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		bookRepo := &mockrepo.BookRepository{}
		bookRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Book{}, nil)
		bookRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BookRepo").Return(bookRepo)

		uc := bookUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllBook(context.Background(), &domain.FilterBook{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		bookRepo := &mockrepo.BookRepository{}
		bookRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Book{}, errors.New("Error"))
		bookRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BookRepo").Return(bookRepo)

		uc := bookUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, _, err := uc.GetAllBook(context.Background(), &domain.FilterBook{})
		assert.Error(t, err)
	})
}
