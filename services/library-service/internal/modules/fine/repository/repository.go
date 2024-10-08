// Code generated by candi v1.17.15.

package repository

import (
	"context"

	"monorepo/services/library-service/internal/modules/fine/domain"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
)

// FineRepository abstract interface
type FineRepository interface {
	FetchAll(ctx context.Context, filter *domain.FilterFine) ([]shareddomain.Fine, error)
	Count(ctx context.Context, filter *domain.FilterFine) int
	Find(ctx context.Context, filter *domain.FilterFine) (shareddomain.Fine, error)
	Save(ctx context.Context, data *shareddomain.Fine, updateOptions ...candishared.DBUpdateOptionFunc) error
	Delete(ctx context.Context, filter *domain.FilterFine) (err error)
}
