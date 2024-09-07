package usecase

import (
	"context"

	"monorepo/services/user-service/internal/modules/user/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *userUsecaseImpl) GetAllUser(ctx context.Context, filter *domain.FilterUser) (result domain.ResponseUserList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:GetAllUser")
	defer trace.Finish()

	data, err := uc.repoSQL.UserRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.UserRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseUser, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail, "")
	}

	return
}
