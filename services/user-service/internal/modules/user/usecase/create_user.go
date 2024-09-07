package usecase

import (
	"context"

	"monorepo/services/user-service/internal/modules/user/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *userUsecaseImpl) CreateUser(ctx context.Context, req *domain.RequestUser) (result domain.ResponseUser, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:CreateUser")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.UserRepo().Save(ctx, &data)
	result.Serialize(&data, "")

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
