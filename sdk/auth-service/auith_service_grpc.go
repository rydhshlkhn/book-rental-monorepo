package authservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	proto "monorepo/sdk/auth-service/proto/token"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/logger"
	"github.com/golangid/candi/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type authserviceGRPCImpl struct {
	host    string
	authKey string
	client  proto.TokenHandlerClient
}

// NewAuthserviceServiceGRPC constructor
func NewAuthserviceServiceGRPC(host string, authKey string) Authservice {

	if u, _ := url.Parse(host); u.Host != "" {
		host = u.Host
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  50 * time.Millisecond,
			Multiplier: 5,
			MaxDelay:   50 * time.Millisecond,
		},
		MinConnectTimeout: 1 * time.Second,
	}))
	if err != nil {
		panic(err)
	}

	return &authserviceGRPCImpl{
		host:    host,
		authKey: authKey,
		client:  proto.NewTokenHandlerClient(conn),
	}
}

func (us *authserviceGRPCImpl) GenerateToken(ctx context.Context, req PayloadGenerateToken) (result ResponseGenerateToken, code int, err error) {
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

func (us *authserviceGRPCImpl) ValidateToken(ctx context.Context, tokenStrnig string) (res *Response, code int, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "AuthServiceSDK:GenerateToken")
	defer trace.Finish()

	headers := map[string]string{
		"Authorization": us.authKey,
	}
	trace.InjectRequestHeader(headers)

	md := metadata.Pairs("authorization", us.authKey)
	ctx = metadata.NewOutgoingContext(ctx, md)
	reqData := &proto.PayloadValidate{
		Token: tokenStrnig,
	}

	trace.SetTag("host", us.host)
	tracer.Log(ctx, "request.data", reqData)

	resp, err := us.client.ValidateToken(ctx, reqData)
	if err != nil {
		trace.SetError(err)
		logger.LogE(err.Error())
		desc, ok := status.FromError(err)
		if ok {
			err = errors.New(desc.Message())
		}
		return
	}

	code = 200

	tracer.Log(ctx, "response.data", resp)

	res = new(Response)
	res.Success = true
	res.Message = "success"
	res.Code = 200
	res.Data = &ResponseClaim{
		StandardClaims: jwt.StandardClaims{
			Audience:  resp.Claim.Audience,
			Subject:   resp.Claim.Subject,
			ExpiresAt: resp.Claim.ExpiresAt,
			Issuer:    resp.Claim.Issuer,
			IssuedAt:  resp.Claim.IssuedAt,
			NotBefore: resp.Claim.NotBefore,
		},
		// User: res.Data.User,
	}

	tracer.Log(ctx, "response1", res)

	return
}
