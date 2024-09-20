package usecase

import (
	"context"
	"errors"
	"monorepo/services/auth-service/internal/modules/token/domain"
	"monorepo/services/auth-service/pkg/shared/domain/constant"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/tracer"
)

var (
	// ErrTokenFormat var
	ErrTokenFormat = errors.New("invalid token format")
	// ErrTokenExpired var
	ErrTokenExpired = errors.New("token is expired")
)

func (uc *tokenUsecaseImpl) Validate(ctx context.Context, tokenString string) (claim *domain.Claim, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenUsecase:Validate")
	defer trace.Finish()

	tokenParse, err := jwt.ParseWithClaims(tokenString, &domain.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.TokenKey), nil
	})

	switch ve := err.(type) {
	case *jwt.ValidationError:
		if ve.Errors == jwt.ValidationErrorExpired {
			err = ErrTokenExpired
		} else {
			err = ErrTokenFormat
		}
	}

	if err != nil {
		return
	}

	if !tokenParse.Valid {
		return claim, ErrTokenFormat
	}

	claim, _ = tokenParse.Claims.(*domain.Claim)

	repoFiilter := domain.FilterToken{DeviceID: claim.DeviceID, UserID: claim.User.ID}
	userToken, err := uc.repoSQL.TokenRepo().Find(ctx, &repoFiilter)
	if err != nil {
		return nil, ErrTokenExpired
	}
	if userToken.IsActive == nil || (userToken.IsActive != nil && !*userToken.IsActive) {
		return nil, ErrTokenExpired
	}
	return
}
