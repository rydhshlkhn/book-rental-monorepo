package authservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
)

type authserviceRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewAuthserviceServiceREST constructor
func NewAuthserviceServiceREST(host string, authKey string) Authservice {

	return &authserviceRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(100000*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("authservice"),
		),
	}
}

func (us *authserviceRESTImpl) GenerateToken(ctx context.Context, req PayloadGenerateToken) (result ResponseGenerateToken, code int, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "AuthServiceSDK:GenerateToken")
	defer trace.Finish()

	headers := map[string]string{
		candihelper.HeaderContentType:   candihelper.HeaderMIMEApplicationJSON,
		candihelper.HeaderAuthorization: fmt.Sprintf("Basic %s", us.authKey),
	}

	reqBody := map[string]string{
		"user_id":   req.UserID,
		"device_id": req.DeviceID,
		"role":      req.Role,
		"username":  req.Username,
	}
	uri := fmt.Sprintf("http://localhost:8001/v1/token/generate")
	body, statusCode, err := candiutils.NewHTTPRequest(
		candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
	).Do(ctx, http.MethodPost, uri, candihelper.ToBytes(reqBody), headers)
	if err != nil {
		trace.SetError(err)
		return result, code, err
	}

	var response struct {
		Success bool                  `json:"success"`
		Data    ResponseGenerateToken `json:"data"`
		Message string                `json:"message"`
	}
	json.Unmarshal(body, &response)
	return response.Data, statusCode, nil
}

func (us *authserviceRESTImpl) ValidateToken(ctx context.Context, tokenStrnig string) (res *Response, code int, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "AuthServiceSDK:GenerateToken")
	defer trace.Finish()

	headers := map[string]string{
		candihelper.HeaderContentType:   candihelper.HeaderMIMEApplicationJSON,
		candihelper.HeaderAuthorization: us.authKey,
	}

	uri := fmt.Sprintf("http://localhost:8001/v1/token/validate?token=%s", tokenStrnig)
	body, code, err := candiutils.NewHTTPRequest(
		candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
	).Do(ctx, http.MethodGet, uri, nil, headers)
	// if err != nil {
	// 	trace.SetError(err)
	// 	return claim, code, err
	// }

	json.Unmarshal(body, &res)
	if err != nil {
		err = errors.New(res.Message)
	}
	return 
}
