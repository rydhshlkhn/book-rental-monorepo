package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"

	"monorepo/services/user-service/internal/modules/user/domain"
	"monorepo/services/user-service/pkg/helper"
	shareddomain "monorepo/services/user-service/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *userUsecaseImpl) Register(ctx context.Context, req *domain.RequestUser) (result domain.ResponseUser, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:Register")
	defer trace.Finish()

	repoFilter := domain.FilterUser{Email: req.Email}
	_, err = uc.repoSQL.UserRepo().Find(ctx, &repoFilter)
	if err == nil {
		err = fmt.Errorf("email %s  already taken", req.Email)
		return
	}

	passHasher := helper.NewPassword(sha1.New, 8, 32, 15000)
	pass := passHasher.HashPassword(req.Password)

	user := shareddomain.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: pass.CipherText,
		PasswordSalt: pass.Salt,
		RoleID:       1,
		StatusID:     1,
	}

	err = uc.repoSQL.UserRepo().Save(ctx, &user)
	result.Serialize(&user, "")

	// Sample using broker publisher
	// uc.deps.GetBroker(types.Kafka). // get registered broker type (sample Kafka)
	// 				GetPublisher().
	// 				PublishMessage(ctx, &candishared.PublisherArgument{
	// 		Topic:   "[topic]",
	// 		Key:     "[key]",
	// 		Message: candihelper.ToBytes("[message]"),
	// 	})
	return
}
