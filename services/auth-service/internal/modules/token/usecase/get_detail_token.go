package usecase

import (
	"context"

	"monorepo/services/auth-service/internal/modules/token/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *tokenUsecaseImpl) GetDetailToken(ctx context.Context, id int) (result domain.ResponseToken, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenUsecase:GetDetailToken")
	defer trace.Finish()

	repoFilter := domain.FilterToken{ID: &id}
	data, err := uc.repoSQL.TokenRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
