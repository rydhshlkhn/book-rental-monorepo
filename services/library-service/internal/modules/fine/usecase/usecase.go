// Code generated by candi v1.17.15.

package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/fine/domain"
	"monorepo/services/library-service/pkg/shared/repository"
	"monorepo/services/library-service/pkg/shared/usecase/common"

	"github.com/golangid/candi/codebase/factory/dependency"
)

// FineUsecase abstraction
type FineUsecase interface {
	GetAllFine(ctx context.Context, filter *domain.FilterFine) (data domain.ResponseFineList, err error)
	GetDetailFine(ctx context.Context, id int) (data domain.ResponseFine, err error)
	CreateFine(ctx context.Context, data *domain.RequestFine) (res domain.ResponseFine, err error)
	UpdateFine(ctx context.Context, data *domain.RequestFine) (err error)
	DeleteFine(ctx context.Context, id int) (err error)
}

type fineUsecaseImpl struct {
	deps          dependency.Dependency
	sharedUsecase common.Usecase
	repoSQL       repository.RepoSQL
	// repoMongo     repository.RepoMongo
}

// NewFineUsecase usecase impl constructor
func NewFineUsecase(deps dependency.Dependency) (FineUsecase, func(sharedUsecase common.Usecase)) {
	uc := &fineUsecaseImpl{
		deps:    deps,
		repoSQL: repository.GetSharedRepoSQL(),
		// repoMongo: repository.GetSharedRepoMongo(),

	}
	return uc, func(sharedUsecase common.Usecase) {
		uc.sharedUsecase = sharedUsecase
	}
}
