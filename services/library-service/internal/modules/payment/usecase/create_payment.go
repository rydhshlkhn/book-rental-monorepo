package usecase

import (
	"context"
	"fmt"

	"monorepo/services/library-service/internal/modules/payment/domain"
	"monorepo/services/library-service/pkg/helper"
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"monorepo/services/library-service/pkg/shared/domain/constant"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *paymentUsecaseImpl) CreatePayment(ctx context.Context, req *domain.RequestTransaction) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:CreatePayment")
	defer trace.Finish()

	data, err := req.Deserialize()
	if err != nil {
		return
	}

	id, err := helper.GetIDFromOrderID(data.OrderID, constant.PREFIX_LENDING)
	if err != nil {
		return
	}

	if isValidSignature := helper.VerifyingSignatureKey(data.SignatureKey, data.OrderID, data.StatusCode, data.GrossAmount); !isValidSignature {
		return fmt.Errorf("invalid signature key")
	}

	lendingFilter := shareddomain.LendingParamGet{ID: &id}
	lendingFilter.Preloads = []string{"Fine"}
	existingLending, err := uc.repoSQL.LendingRepo().Find(ctx, &lendingFilter)
	existingFine := existingLending.Fine
	if err != nil {
		return
	}

	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) (err error) {
		existingFine.PaymentStatus = data.TransactionStatus
		if err = uc.repoSQL.FineRepo().Save(ctx, &existingFine, candishared.DBUpdateSetUpdatedFields("PaymentStatus")); err != nil {
			return
		}

		if err = uc.repoSQL.PaymentRepo().Save(ctx, &data); err != nil {
			return
		}

		return nil
	})

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
