package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"
	mockrepo "monorepo/services/library-service/pkg/mocks/modules/book/repository"
	mocksharedrepo "monorepo/services/library-service/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bookUsecaseImpl_CreateBook(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		bookRepo := &mockrepo.BookRepository{}
		bookRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BookRepo").Return(bookRepo)

		uc := bookUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateBook(context.Background(), &domain.RequestBook{})
		assert.NoError(t, err)
	})
}
