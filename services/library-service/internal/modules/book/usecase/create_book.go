package usecase

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) CreateBook(ctx context.Context, req *domain.RequestBook) (result domain.ResponseBook, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:CreateBook")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.BookRepo().SaveBook(ctx, &data)
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
