package usecase

import (
	"context"
	"errors"
	"monorepo/services/auth-service/internal/modules/token/domain"
	shareddomain "monorepo/services/auth-service/pkg/shared/domain"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/tracer"
)

const (
	// TokenKey const
	TokenKey = "18608c7d-b319-0xc000165c80-0xc0000da000-11478e4e2650"
)

var (
	// ErrTokenFormat var
	ErrTokenFormat = errors.New("invalid token format")
	// ErrTokenExpired var
	ErrTokenExpired = errors.New("token is expired")
)

func (uc *tokenUsecaseImpl) Generate(ctx context.Context, payload *domain.Claim) (result domain.ResponseGenerateToken, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenUsecase:Generate")
	defer trace.Finish()

	now := time.Now()

	repoFilter := domain.FilterToken{DeviceID: payload.DeviceID, UserID: payload.User.ID}
	savedToken := shareddomain.Token{UserID: payload.User.ID, DeviceID: payload.DeviceID}
	if savedToken, err = uc.repoSQL.TokenRepo().Find(ctx, &repoFilter); err == nil &&
		candihelper.PtrToBool(savedToken.IsActive) &&
		savedToken.ExpiredAt.After(now) {
		return domain.ResponseGenerateToken{
			Token:        savedToken.Token,
			RefreshToken: savedToken.RefreshToken,
		}, nil
	}

	exp := now.Add(time.Hour * 10)

	key := []byte(TokenKey)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := jwt.MapClaims{
		"iss":  "mooc",
		"exp":  exp.Unix(),
		"iat":  now.Unix(),
		"did":  payload.DeviceID,
		"aud":  payload.Audience,
		"jti":  payload.Id,
		"sub":  payload.User.ID,
		"user": payload.User,
	}
	token.Claims = claims

	tokenString, err := token.SignedString(key)
	if err != nil {
		return
	}

	// refresh token
	refreshTokenHS := jwt.New(jwt.SigningMethodHS256)
	refreshTokenHS.Claims = jwt.MapClaims{
		"exp": exp.Unix(),
	}

	refreshTokenString, err := refreshTokenHS.SignedString(key)
	if err != nil {
		return
	}

	savedToken.Token = tokenString
	savedToken.RefreshToken = refreshTokenString
	savedToken.DeviceID = payload.DeviceID
	savedToken.UserID = payload.User.ID
	savedToken.ExpiredAt = exp
	savedToken.IsActive = candihelper.ToBoolPtr(true)
	uc.repoSQL.TokenRepo().Save(ctx, &savedToken)

	result.Token = tokenString
	result.RefreshToken = refreshTokenString

	return
}

func (uc *tokenUsecaseImpl) Validate(ctx context.Context, tokenString string) (claim *domain.Claim, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "TokenUsecase:Validate")
	defer trace.Finish()

	tokenParse, err := jwt.ParseWithClaims(tokenString, &domain.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenKey), nil
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
