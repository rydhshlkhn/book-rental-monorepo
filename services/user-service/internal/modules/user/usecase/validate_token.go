package usecase

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *userUsecaseImpl) ValidateToken(ctx context.Context, tokenString string) (claim *candishared.TokenClaim, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:ValidateToken")
	defer trace.Finish()

	resp, _, err := uc.sdk.Authservice().ValidateToken(ctx, tokenString)
	if err != nil {
		err = fmt.Errorf("%s", resp.Message)
		return 
	}

	claim = new(candishared.TokenClaim)
	claim.StandardClaims = jwt.StandardClaims{
		Audience:  resp.Data.Audience,
		Subject:   resp.Data.Subject,
		ExpiresAt: resp.Data.ExpiresAt,
		Issuer:    resp.Data.Issuer,
		IssuedAt:  resp.Data.IssuedAt,
		NotBefore: resp.Data.NotBefore,
	}
	claim.Additional = map[string]interface{}{
		"username": resp.Data.User.Username,
		"user_id":  resp.Data.User.ID,
	}

	return
}
