package authservice

import "context"

// Authservice client abstract interface

type Authservice interface {
	// Add service client method
	GenerateToken(ctx context.Context, req PayloadGenerateToken) (result ResponseGenerateToken, code int, err error)
	ValidateToken(ctx context.Context, tokenStrnig string) (claim *ResponseClaim, code int, err error)
}
