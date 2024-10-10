package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/activity/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *activityUsecaseImpl) CreateActivity(ctx context.Context, req *domain.RequestActivity) (result domain.ResponseActivity, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ActivityUsecase:CreateActivity")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.ActivityRepo().Save(ctx, &data)
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
