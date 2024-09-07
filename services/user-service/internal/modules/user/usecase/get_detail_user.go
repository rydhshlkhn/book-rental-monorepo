package usecase

import (
	"context"

	"monorepo/services/user-service/internal/modules/user/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *userUsecaseImpl) GetDetailUser(ctx context.Context, id int) (result domain.ResponseUser, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:GetDetailUser")
	defer trace.Finish()

	repoFilter := domain.FilterUser{ID: &id}
	data, err := uc.repoSQL.UserRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data, "")
	return
}
