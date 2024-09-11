package usecase

import (
	"context"

	mockrepo "monorepo/services/book-service/pkg/mocks/modules/book/repository"
	mocksharedrepo "monorepo/services/book-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/book-service/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bookUsecaseImpl_GetDetailBook(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		bookRepo := &mockrepo.BookRepository{}
		bookRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Book{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BookRepo").Return(bookRepo)

		uc := bookUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailBook(context.Background(), 1)
		assert.NoError(t, err)
	})
}
