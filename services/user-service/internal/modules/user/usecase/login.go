package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"strconv"

	authservice "monorepo/sdk/auth-service"
	"monorepo/services/user-service/internal/modules/user/domain"
	"monorepo/services/user-service/pkg/helper"

	"github.com/golangid/candi/tracer"
)

func (uc *userUsecaseImpl) Login(ctx context.Context, req *domain.RequestLoginUser) (result domain.ResponseUser, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:Login")
	defer trace.Finish()

	repoFilter := domain.FilterUser{Email: req.Email}
	repoFilter.Preloads = []string{"Role"}
	user, err := uc.repoSQL.UserRepo().Find(ctx, &repoFilter)
	if err != nil {
		return
	}

	passHasher := helper.NewPassword(sha1.New, 8, 32, 15000)
	if !passHasher.VerifyPassword(req.Password, user.PasswordHash, user.PasswordSalt) {
		err = errors.New("invalid Password")
		return
	}

	payload := authservice.PayloadGenerateToken{DeviceID: "Web", UserID: strconv.Itoa(user.ID), Role: user.Role.Name, Username: user.Username}
	response, _, err := uc.sdk.Authservice().GenerateToken(ctx, payload)
	if err != nil {
		return
	}

	result.Serialize(&user, response.Token)

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
