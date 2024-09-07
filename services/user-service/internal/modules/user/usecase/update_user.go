package usecase

import (
	"context"

	"monorepo/services/user-service/internal/modules/user/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *userUsecaseImpl) UpdateUser(ctx context.Context, data *domain.RequestUser) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:UpdateUser")
	defer trace.Finish()

	repoFilter := domain.FilterUser{ID: &data.ID}
	existing, err := uc.repoSQL.UserRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Username = data.Username
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.UserRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
