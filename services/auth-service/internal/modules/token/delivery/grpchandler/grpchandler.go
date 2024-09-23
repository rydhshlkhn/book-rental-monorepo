// Code generated by candi v1.17.15.

package grpchandler

import (
	"context"

	proto "monorepo/sdk/auth-service/proto/token"
	"monorepo/services/auth-service/internal/modules/token/domain"
	"monorepo/services/auth-service/pkg/shared/usecase"

	"google.golang.org/grpc"

	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/factory/types"
	"github.com/golangid/candi/codebase/interfaces"
	"github.com/golangid/candi/tracer"
)

// GRPCHandler rpc handler
type GRPCHandler struct {
	mw        interfaces.Middleware
	uc        usecase.Usecase
	validator interfaces.Validator
}

// NewGRPCHandler func
func NewGRPCHandler(uc usecase.Usecase, deps dependency.Dependency) *GRPCHandler {
	return &GRPCHandler{
		uc: uc, mw: deps.GetMiddleware(), validator: deps.GetValidator(),
	}
}

// Register grpc server
func (h *GRPCHandler) Register(server *grpc.Server, mwGroup *types.MiddlewareGroup) {
	proto.RegisterTokenHandlerServer(server, h)

	// register middleware for method
	mwGroup.AddProto(proto.File_token_token_proto, h.ValidateToken, h.mw.GRPCBasicAuth)
	mwGroup.AddProto(proto.File_token_token_proto, h.GenerateToken, h.mw.GRPCBasicAuth)
}

// GetAllToken rpc method
func (h *GRPCHandler) ValidateToken(ctx context.Context, req *proto.PayloadValidate) (*proto.ResponseValidation, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenDeliveryGRPC:ValidateToken")
	defer trace.Finish()

	token := req.Token
	claim, err := h.uc.Token().Validate(ctx, token)
	if err != nil {
		return nil, err
	}

	return &proto.ResponseValidation{
		Success: true,
		Claim: &proto.ResponseValidation_ClaimData{
			Audience:  claim.Audience,
			Subject:   claim.Subject,
			ExpiresAt: claim.ExpiresAt,
			User: &proto.UserData{
				ID:       claim.User.ID,
				Username: claim.User.Username,
			},
		},
	}, nil
}

// GenerateToken rpc method
func (h *GRPCHandler) GenerateToken(ctx context.Context, req *proto.UserData) (*proto.ResponseGenerate, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenDeliveryGRPC:GenerateToken")
	defer trace.Finish()

	var tokenClaim domain.Claim
	tokenClaim.User.ID = req.ID
	tokenClaim.User.Role = req.RoleID
	tokenClaim.User.Username = req.Username
	tokenClaim.DeviceID = req.DeviceID

	data, err := h.uc.Token().Generate(ctx, &tokenClaim)
	if err != nil {
		return nil, err
	}

	return &proto.ResponseGenerate{
		Success: true,
		Data: &proto.ResponseGenerate_Token{
			Token:        data.Token,
			RefreshToken: data.RefreshToken,
		},
	}, nil
}
