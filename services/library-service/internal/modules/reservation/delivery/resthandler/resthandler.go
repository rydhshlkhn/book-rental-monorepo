// Code generated by candi v1.17.15.

package resthandler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"monorepo/services/library-service/internal/modules/reservation/domain"
	"monorepo/services/library-service/pkg/shared/usecase"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	restserver "github.com/golangid/candi/codebase/app/rest_server"
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
	v1Reservation := root.Group(candihelper.V1+"/reservation", h.mw.HTTPBearerAuth)

	v1Reservation.GET("/", h.getAllReservation, h.mw.HTTPPermissionACL("getAllReservation"))
	v1Reservation.GET("/:id", h.getDetailReservationByID, h.mw.HTTPPermissionACL("getDetailReservation"))
	v1Reservation.POST("/", h.createReservation, h.mw.HTTPPermissionACL("createReservation"))
	v1Reservation.PUT("/:id", h.updateReservation, h.mw.HTTPPermissionACL("updateReservation"))
	v1Reservation.DELETE("/:id", h.deleteReservation, h.mw.HTTPPermissionACL("deleteReservation"))
}

// GetAllReservation documentation
// @Summary			Get All Reservation
// @Description		API for get all reservation
// @Tags			Reservation
// @Accept			json
// @Produce			json
// @Param			page	query	string	false	"Page with default value is 1"
// @Param			limit	query	string	false	"Limit with default value is 10"
// @Param			search	query	string	false	"Search"
// @Param			orderBy	query	string	false	"Order By"
// @Param			sort	query	string	false	"Sort (ASC DESC)"
// @Success			200	{object}	domain.ResponseReservationList
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/reservation [get]
func (h *RestHandler) getAllReservation(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ReservationDeliveryREST:GetAllReservation")
	defer trace.Finish()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using HTTPBearerAuth in middleware for this handler

	var filter domain.FilterReservation
	if err := candihelper.ParseFromQueryParam(req.URL.Query(), &filter); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed parse filter", err).JSON(rw)
		return
	}

	if err := h.validator.ValidateDocument("reservation/get_all", filter); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate filter", err).JSON(rw)
		return
	}

	result, err := h.uc.Reservation().GetAllReservation(ctx, &filter)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	message := "Success, with your user id (" + tokenClaim.Subject + ") and role (" + tokenClaim.Role + ")"
	wrapper.NewHTTPResponse(http.StatusOK, message, result.Meta, result.Data).JSON(rw)
}

// GetDetailReservation documentation
// @Summary			Get Detail Reservation
// @Description		API for get detail reservation
// @Tags			Reservation
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Success			200	{object}	domain.ResponseReservation
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/reservation/{id} [get]
func (h *RestHandler) getDetailReservationByID(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ReservationDeliveryREST:GetDetailReservationByID")
	defer trace.Finish()

	id, _ := strconv.Atoi(restserver.URLParam(req, "id"))
	data, err := h.uc.Reservation().GetDetailReservation(ctx, id)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success", data).JSON(rw)
}

// CreateReservation documentation
// @Summary			Create Reservation
// @Description		API for create reservation
// @Tags			Reservation
// @Accept			json
// @Produce			json
// @Param			data	body	domain.RequestReservation	true	"Body Data"
// @Success			200	{object}	domain.ResponseReservation
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/reservation [post]
func (h *RestHandler) createReservation(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ReservationDeliveryREST:CreateReservation")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("reservation/save", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.RequestReservation
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	res, err := h.uc.Reservation().CreateReservation(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusCreated, "Success", res).JSON(rw)
}

// UpdateReservation documentation
// @Summary			Update Reservation
// @Description		API for update reservation
// @Tags			Reservation
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Param			data	body	domain.RequestReservation	true	"Body Data"
// @Success			200	{object}	domain.ResponseReservation
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/reservation/{id} [put]
func (h *RestHandler) updateReservation(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ReservationDeliveryREST:UpdateReservation")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("reservation/save", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.RequestReservation
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	payload.ID, _ = strconv.Atoi(restserver.URLParam(req, "id"))
	err := h.uc.Reservation().UpdateReservation(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(rw)
}

// DeleteReservation documentation
// @Summary			Delete Reservation
// @Description		API for delete reservation
// @Tags			Reservation
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Success			200	{object}	domain.ResponseReservation
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/reservation/{id} [delete]
func (h *RestHandler) deleteReservation(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ReservationDeliveryREST:DeleteReservation")
	defer trace.Finish()

	id, _ := strconv.Atoi(restserver.URLParam(req, "id"))
	if err := h.uc.Reservation().DeleteReservation(ctx, id); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(rw)
}
