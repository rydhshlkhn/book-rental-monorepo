package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/fine/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *fineUsecaseImpl) CreateFine(ctx context.Context, req *domain.RequestFine) (result domain.ResponseFine, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "FineUsecase:CreateFine")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.FineRepo().Save(ctx, &data)
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
