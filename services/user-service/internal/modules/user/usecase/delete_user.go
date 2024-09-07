package usecase

import (
	"context"
	
	"monorepo/services/user-service/internal/modules/user/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *userUsecaseImpl) DeleteUser(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:DeleteUser")
	defer trace.Finish()

	repoFilter := domain.FilterUser{ID: &id}
	return uc.repoSQL.UserRepo().Delete(ctx, &repoFilter)
}
