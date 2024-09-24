// Code generated by candi v1.17.15.

package usecase

import (
	"context"

	"monorepo/sdk"
	"monorepo/services/library-service/internal/modules/book/domain"
	"monorepo/services/library-service/pkg/shared/repository"
	"monorepo/services/library-service/pkg/shared/usecase/common"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/codebase/factory/dependency"
)

// BookUsecase abstraction
type BookUsecase interface {
	Asdf(ctx context.Context, req *domain.RequestAsdf) (resp domain.ResponseAsdf, err error)
	GetAllBook(ctx context.Context, filter *domain.FilterBook) (data domain.ResponseBookList, err error)
	GetDetailBook(ctx context.Context, id int) (data domain.ResponseBook, err error)
	CreateBook(ctx context.Context, data *domain.RequestBook) (res domain.ResponseBook, err error)
	UpdateBook(ctx context.Context, data *domain.RequestBook) (err error)
	DeleteBook(ctx context.Context, id int) (err error)
	//Book Item
	GetDetailBookItem(ctx context.Context, id int) (result domain.ResponseBookItem, err error)
	CreateBookItem(ctx context.Context, req *domain.RequestBookItem) (result domain.ResponseBookItem, err error)
	UpdateBookItem(ctx context.Context, data *domain.RequestBookItem) (err error)
	DeleteBookItem(ctx context.Context, id int) (err error)
	// middleware
	ValidateToken(ctx context.Context, tokenString string) (claim *candishared.TokenClaim, err error)
}

type bookUsecaseImpl struct {
	deps          dependency.Dependency
	sharedUsecase common.Usecase
	repoSQL       repository.RepoSQL
	sdk           sdk.SDK
	// repoMongo     repository.RepoMongo
}

// NewBookUsecase usecase impl constructor
func NewBookUsecase(deps dependency.Dependency) (BookUsecase, func(sharedUsecase common.Usecase)) {
	uc := &bookUsecaseImpl{
		deps:    deps,
		repoSQL: repository.GetSharedRepoSQL(),
		sdk:     sdk.GetSDK(),
		// repoMongo: repository.GetSharedRepoMongo(),

	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}
