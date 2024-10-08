// Code generated by candi v1.17.15.

package resthandler

import (
	"encoding/json"
	"io"
	"net/http"

	"monorepo/services/library-service/internal/modules/payment/domain"
	"monorepo/services/library-service/pkg/shared/usecase"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/interfaces"
	"github.com/golangid/candi/tracer"
	"github.com/golangid/candi/wrapper"
)

// RestHandler handler
type RestHandler struct {
	mw        interfaces.Middleware
	uc        usecase.Usecase
	validator interfaces.Validator
}

// NewRestHandler create new rest handler
func NewRestHandler(uc usecase.Usecase, deps dependency.Dependency) *RestHandler {
	return &RestHandler{
		uc: uc, mw: deps.GetMiddleware(), validator: deps.GetValidator(),
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root interfaces.RESTRouter) {
	v1Payment := root.Group(candihelper.V1 + "/payment")

	v1Payment.POST("/midtrans-notifiicaton", h.miidtransNotification)
}

// CreatePayment documentation
// @Summary			Create Payment
// @Description		API for create payment
// @Tags			Payment
// @Accept			json
// @Produce			json
// @Param			data	body	domain.RequestPayment	true	"Body Data"
// @Success			200	{object}	domain.ResponsePayment
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/payment [post]
func (h *RestHandler) miidtransNotification(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "PaymentDeliveryREST:CreatePayment")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("payment/save", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.RequestTransaction
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	err := h.uc.Payment().CreatePayment(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusCreated, "Success").JSON(rw)
}
