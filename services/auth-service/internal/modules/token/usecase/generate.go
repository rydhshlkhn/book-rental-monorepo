package usecase

import (
	"context"
	"monorepo/services/auth-service/internal/modules/token/domain"
	"monorepo/services/auth-service/pkg/shared"
	shareddomain "monorepo/services/auth-service/pkg/shared/domain"
	constant "monorepo/services/auth-service/pkg/shared/domain/constant"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
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
	ageDuration := shared.GetEnv().JWTAccessTokenAge
	exp := now.Add(ageDuration)

	key := []byte(constant.TokenKey)
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

	uc.redisPub.PublishMessage(ctx, &candishared.PublisherArgument{
		Topic:   domain.RedisTokenExpiredKeyConst,
		Key:     domain.RedisTokenExpiredKeyConst,
		Message: candihelper.ToBytes(claims),
		Delay:   ageDuration,
	})

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
