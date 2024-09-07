package usecase

import (
	"context"

	"monorepo/services/auth-service/internal/modules/token/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *tokenUsecaseImpl) CreateToken(ctx context.Context, req *domain.RequestToken) (result domain.ResponseToken, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenUsecase:CreateToken")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.TokenRepo().Save(ctx, &data)
	result.Serialize(&data)

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
