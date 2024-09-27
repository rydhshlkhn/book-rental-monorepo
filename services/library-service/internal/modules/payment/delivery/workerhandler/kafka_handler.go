// Code generated by candi v1.17.15.

package workerhandler

import (
	"fmt"

	"monorepo/services/library-service/pkg/shared/usecase"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/factory/types"
	"github.com/golangid/candi/codebase/interfaces"
	"github.com/golangid/candi/tracer"
)

// KafkaHandler struct
type KafkaHandler struct {
	uc        usecase.Usecase
	validator interfaces.Validator
}

// NewKafkaHandler constructor
func NewKafkaHandler(uc usecase.Usecase, deps dependency.Dependency) *KafkaHandler {
	return &KafkaHandler{
		uc:        uc,
		validator: deps.GetValidator(),
	}
}

// MountHandlers mount handler group
func (h *KafkaHandler) MountHandlers(group *types.WorkerHandlerGroup) {
	group.Add("lending", h.handlePayment) // handling topic "payment"
}

// ProcessMessage from kafka consumer
func (h *KafkaHandler) handlePayment(eventContext *candishared.EventContext) error {
	trace, ctx := tracer.StartTraceWithContext(eventContext.Context(), "PaymentDeliveryKafka:HandlePayment")
	defer trace.Finish()

	fmt.Printf("message consumed in handler %s. key: %s, message: %s\n", eventContext.HandlerRoute(), eventContext.Key(), eventContext.Message())

	// exec usecase
	// h.uc.SomethingUsecase()

	return ctx.Err()
}
