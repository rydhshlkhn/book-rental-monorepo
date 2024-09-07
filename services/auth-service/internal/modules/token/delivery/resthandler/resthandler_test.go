// Code generated by candi v1.17.15.

package resthandler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"monorepo/services/auth-service/internal/modules/token/domain"
	mockusecase "monorepo/services/auth-service/pkg/mocks/modules/token/usecase"
	mocksharedusecase "monorepo/services/auth-service/pkg/mocks/shared/usecase"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	mockdeps "github.com/golangid/candi/mocks/codebase/factory/dependency"
	mockinterfaces "github.com/golangid/candi/mocks/codebase/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, reqBody                       string
	wantValidateError, wantUsecaseError error
	wantRespCode                        int
}

var (
	errFoo = errors.New("Something error")
)

func TestNewRestHandler(t *testing.T) {
	mockMiddleware := &mockinterfaces.Middleware{}
	mockMiddleware.On("HTTPPermissionACL", mock.Anything).Return(func(http.Handler) http.Handler { return nil })
	mockValidator := &mockinterfaces.Validator{}

	mockDeps := &mockdeps.Dependency{}
	mockDeps.On("GetMiddleware").Return(mockMiddleware)
	mockDeps.On("GetValidator").Return(mockValidator)

	handler := NewRestHandler(nil, mockDeps)
	assert.NotNil(t, handler)

	mockRoute := &mockinterfaces.RESTRouter{}
	mockRoute.On("Group", mock.Anything, mock.Anything).Return(mockRoute)
	mockRoute.On("GET", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("POST", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("PUT", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("DELETE", mock.Anything, mock.Anything, mock.Anything)
	handler.Mount(mockRoute)
}

func TestRestHandler_generateToken(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: `{"user_id": "test@test.com", "device_id": "Web"}`, wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", reqBody: `{"user_id": "test@test.com", "device_id": "Web"}`, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tokenUsecase := &mockusecase.TokenUsecase{}
			tokenUsecase.On("Generate", mock.Anything, mock.Anything).Return(domain.ResponseGenerateToken{}, tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("Token").Return(tokenUsecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(candihelper.HeaderContentType, candihelper.HeaderMIMEApplicationJSON)
			res := httptest.NewRecorder()
			handler.generateToken(res, req)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}

func TestRestHandler_validateToken(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: "?token=str", wantUsecaseError: nil, wantRespCode: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tokenUsecase := &mockusecase.TokenUsecase{}
			tokenUsecase.On("Validate", mock.Anything, mock.Anything).Return(
				&domain.Claim{}, tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("Token").Return(tokenUsecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodGet, "/"+tt.reqBody, strings.NewReader(tt.reqBody))
			req = req.WithContext(candishared.SetToContext(req.Context(), candishared.ContextKeyTokenClaim, &candishared.TokenClaim{}))
			req.Header.Add(candihelper.HeaderContentType, candihelper.HeaderMIMEApplicationJSON)
			res := httptest.NewRecorder()
			handler.validateToken(res, req)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}
