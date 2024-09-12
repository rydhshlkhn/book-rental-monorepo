// Code generated by candi v1.17.15.

package repository

import (
	"context"

	"monorepo/services/book-service/internal/modules/book/domain"
	shareddomain "monorepo/services/book-service/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
)

// BookRepository abstract interface
type BookRepository interface {
	FetchAll(ctx context.Context, filter *domain.FilterBook) ([]shareddomain.Book, error)
	Count(ctx context.Context, filter *domain.FilterBook) int
	Find(ctx context.Context, filter *domain.FilterBook) (shareddomain.Book, error)
	SaveBook(ctx context.Context, data *shareddomain.Book, updateOptions ...candishared.DBUpdateOptionFunc) error
	Delete(ctx context.Context, filter *domain.FilterBook) (err error)
	//Book Item
	FindItem(ctx context.Context, filter *domain.FilterBookItem) (result shareddomain.BookItem, err error)
	SaveBookItem(ctx context.Context, data *shareddomain.BookItem, updateOptions ...candishared.DBUpdateOptionFunc) (err error)
	DeleteItem(ctx context.Context, filter *domain.FilterBookItem) (err error)
}
