package usecase

import (
	"context"
	
	"monorepo/services/auth-service/internal/modules/token/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *tokenUsecaseImpl) DeleteToken(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenUsecase:DeleteToken")
	defer trace.Finish()

	repoFilter := domain.FilterToken{ID: &id}
	return uc.repoSQL.TokenRepo().Delete(ctx, &repoFilter)
}
