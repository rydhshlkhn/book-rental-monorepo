package usecase

import (
	"context"

	"monorepo/services/auth-service/internal/modules/token/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *tokenUsecaseImpl) UpdateToken(ctx context.Context, data *domain.RequestToken) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenUsecase:UpdateToken")
	defer trace.Finish()

	repoFilter := domain.FilterToken{ID: &data.ID}
	existing, err := uc.repoSQL.TokenRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Field = data.Field
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.TokenRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
