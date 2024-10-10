// Code generated by candi v1.17.15.

package resthandler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"monorepo/services/library-service/internal/modules/activity/domain"
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
	v1Activity := root.Group(candihelper.V1+"/activity", h.mw.HTTPBearerAuth)

	v1Activity.GET("/", h.getAllActivity, h.mw.HTTPPermissionACL("getAllActivity"))
	v1Activity.GET("/:id", h.getDetailActivityByID, h.mw.HTTPPermissionACL("getDetailActivity"))
	v1Activity.POST("/", h.createActivity, h.mw.HTTPPermissionACL("createActivity"))
	v1Activity.PUT("/:id", h.updateActivity, h.mw.HTTPPermissionACL("updateActivity"))
	v1Activity.DELETE("/:id", h.deleteActivity, h.mw.HTTPPermissionACL("deleteActivity"))
}

// GetAllActivity documentation
// @Summary			Get All Activity
// @Description		API for get all activity
// @Tags			Activity
// @Accept			json
// @Produce			json
// @Param			page	query	string	false	"Page with default value is 1"
// @Param			limit	query	string	false	"Limit with default value is 10"
// @Param			search	query	string	false	"Search"
// @Param			orderBy	query	string	false	"Order By"
// @Param			sort	query	string	false	"Sort (ASC DESC)"
// @Success			200	{object}	domain.ResponseActivityList
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/activity [get]
func (h *RestHandler) getAllActivity(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ActivityDeliveryREST:GetAllActivity")
	defer trace.Finish()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using HTTPBearerAuth in middleware for this handler

	var filter domain.FilterActivity
	if err := candihelper.ParseFromQueryParam(req.URL.Query(), &filter); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed parse filter", err).JSON(rw)
		return
	}

	if err := h.validator.ValidateDocument("activity/get_all", filter); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate filter", err).JSON(rw)
		return
	}

	result, err := h.uc.Activity().GetAllActivity(ctx, &filter)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	message := "Success, with your user id (" + tokenClaim.Subject + ") and role (" + tokenClaim.Role + ")"
	wrapper.NewHTTPResponse(http.StatusOK, message, result.Meta, result.Data).JSON(rw)
}

// GetDetailActivity documentation
// @Summary			Get Detail Activity
// @Description		API for get detail activity
// @Tags			Activity
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Success			200	{object}	domain.ResponseActivity
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/activity/{id} [get]
func (h *RestHandler) getDetailActivityByID(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ActivityDeliveryREST:GetDetailActivityByID")
	defer trace.Finish()

	id, _ := strconv.Atoi(restserver.URLParam(req, "id"))
	data, err := h.uc.Activity().GetDetailActivity(ctx, id)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success", data).JSON(rw)
}

// CreateActivity documentation
// @Summary			Create Activity
// @Description		API for create activity
// @Tags			Activity
// @Accept			json
// @Produce			json
// @Param			data	body	domain.RequestActivity	true	"Body Data"
// @Success			200	{object}	domain.ResponseActivity
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/activity [post]
func (h *RestHandler) createActivity(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ActivityDeliveryREST:CreateActivity")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("activity/save", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.RequestActivity
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	res, err := h.uc.Activity().CreateActivity(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusCreated, "Success", res).JSON(rw)
}

// UpdateActivity documentation
// @Summary			Update Activity
// @Description		API for update activity
// @Tags			Activity
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Param			data	body	domain.RequestActivity	true	"Body Data"
// @Success			200	{object}	domain.ResponseActivity
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/activity/{id} [put]
func (h *RestHandler) updateActivity(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ActivityDeliveryREST:UpdateActivity")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("activity/save", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.RequestActivity
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	payload.ID, _ = strconv.Atoi(restserver.URLParam(req, "id"))
	err := h.uc.Activity().UpdateActivity(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(rw)
}

// DeleteActivity documentation
// @Summary			Delete Activity
// @Description		API for delete activity
// @Tags			Activity
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Success			200	{object}	domain.ResponseActivity
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/activity/{id} [delete]
func (h *RestHandler) deleteActivity(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "ActivityDeliveryREST:DeleteActivity")
	defer trace.Finish()
	
	id, _ := strconv.Atoi(restserver.URLParam(req, "id"))
	if err := h.uc.Activity().DeleteActivity(ctx, id); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success").JSON(rw)
}
