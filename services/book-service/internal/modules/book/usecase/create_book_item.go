package usecase

import (
	"context"

	"monorepo/services/book-service/internal/modules/book/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) CreateBookItem(ctx context.Context, req *domain.RequestBookItem) (result domain.ResponseBookItem, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:CreateBookItem")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.BookRepo().SaveBookItem(ctx, &data)
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
