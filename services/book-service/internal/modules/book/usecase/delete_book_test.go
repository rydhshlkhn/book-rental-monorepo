package usecase

import (
	"context"

	mockrepo "monorepo/services/book-service/pkg/mocks/modules/book/repository"
	mocksharedrepo "monorepo/services/book-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bookUsecaseImpl_DeleteBook(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		bookRepo := &mockrepo.BookRepository{}
		bookRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BookRepo").Return(bookRepo)

		uc := bookUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteBook(context.Background(), 1)
		assert.NoError(t, err)
	})
}
