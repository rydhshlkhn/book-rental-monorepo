package usecase

import (
	"context"

	"monorepo/services/auth-service/internal/modules/token/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *tokenUsecaseImpl) GetAllToken(ctx context.Context, filter *domain.FilterToken) (result domain.ResponseTokenList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenUsecase:GetAllToken")
	defer trace.Finish()

	data, err := uc.repoSQL.TokenRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.TokenRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseToken, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
