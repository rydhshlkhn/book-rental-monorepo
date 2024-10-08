// Code generated by candi v1.17.15.

package resthandler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"monorepo/services/user-service/internal/modules/user/domain"
	"monorepo/services/user-service/pkg/shared/usecase"

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
	v1User := root.Group(candihelper.V1 + "/user")

	v1User.POST("/register", h.registerUser)
	v1User.POST("/login", h.loginUser)
	v1User.GET("/:id", h.getDetailUserByID, h.mw.HTTPBearerAuth, h.mw.HTTPPermissionACL("getDetailUser"))
}

// RegisterUser documentation
// @Summary			Register User
// @Description		API for create user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			data	body	domain.RequestUser	true	"Body Data"
// @Success			200	{object}	domain.ResponseUser
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/user/register [post]
func (h *RestHandler) registerUser(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "UserDeliveryREST:CreateUser")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("user/register", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.RequestUser
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	res, err := h.uc.User().Register(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusCreated, "Success", res).JSON(rw)
}

// LoginUser documentation
// @Summary			Login User
// @Description		API for create user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			data	body	domain.RequestUser	true	"Body Data"
// @Success			200	{object}	domain.ResponseUser
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/user/login [post]
func (h *RestHandler) loginUser(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "UserDeliveryREST:CreateUser")
	defer trace.Finish()

	body, _ := io.ReadAll(req.Body)
	if err := h.validator.ValidateDocument("user/login", body); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate payload", err).JSON(rw)
		return
	}

	var payload domain.RequestLoginUser
	if err := json.Unmarshal(body, &payload); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	res, err := h.uc.User().Login(ctx, &payload)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusCreated, "Success", res).JSON(rw)
}

// GetDetailUser documentation
// @Summary			Get Detail User
// @Description		API for get detail user
// @Tags			User
// @Accept			json
// @Produce			json
// @Param			id	path	string	true	"ID"
// @Success			200	{object}	domain.ResponseUser
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/user/{id} [get]
func (h *RestHandler) getDetailUserByID(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "UserDeliveryREST:GetDetailUserByID")
	defer trace.Finish()

	tokenClaim := candishared.ParseTokenClaimFromContext(ctx)
	if tokenClaim.Subject != restserver.URLParam(req, "id") {
		err := errors.New("not Allowed")
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	id, _ := strconv.Atoi(restserver.URLParam(req, "id"))
	data, err := h.uc.User().GetDetailUser(ctx, id)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	wrapper.NewHTTPResponse(http.StatusOK, "Success", data).JSON(rw)
}
