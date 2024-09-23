package usecase

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) ValidateToken(ctx context.Context, tokenString string) (claim *candishared.TokenClaim, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:ValidateToken")
	defer trace.Finish()

	resp, _, err := uc.sdk.Authservice().ValidateToken(ctx, tokenString)
	if err != nil {
		// err = fmt.Errorf("%s", resp.Message)
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

	// tokenParse, err := jwt.ParseWithClaims(tokenString, &candishared.TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(TokenKey), nil
	// })

	// switch ve := err.(type) {
	// case *jwt.ValidationError:
	// 	if ve.Errors == jwt.ValidationErrorExpired {
	// 		err = ErrTokenExpired
	// 	} else {
	// 		err = ErrTokenFormat
	// 	}
	// }

	// if err != nil {
	// 	return
	// }

	// if !tokenParse.Valid {
	// 	return claim, ErrTokenFormat
	// }

	// claim, _ = tokenParse.Claims.(*candishared.TokenClaim)

	return
}
