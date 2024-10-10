// Code generated by candi v1.17.15.

package repository

import (
	"context"

	"monorepo/services/library-service/internal/modules/activity/domain"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
)

// ActivityRepository abstract interface
type ActivityRepository interface {
	FetchAll(ctx context.Context, filter *domain.FilterActivity) ([]shareddomain.Activity, error)
	Count(ctx context.Context, filter *domain.FilterActivity) int
	Find(ctx context.Context, filter *domain.FilterActivity) (shareddomain.Activity, error)
	Save(ctx context.Context, data *shareddomain.Activity, updateOptions ...candishared.DBUpdateOptionFunc) error
	Delete(ctx context.Context, filter *domain.FilterActivity) (err error)
}
