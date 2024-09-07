package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"strconv"

	authservice "monorepo/sdk/auth-service"
	"monorepo/services/user-service/internal/modules/user/domain"
	"monorepo/services/user-service/pkg/helper"
	shareddomain "monorepo/services/user-service/pkg/shared/domain"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

const (
	TokenKey = "18608c7d-b319-0xc000165c80-0xc0000da000-11478e4e2650"
)

var (
	// ErrTokenFormat var
	ErrTokenFormat = errors.New("invalid token format")
	// ErrTokenExpired var
	ErrTokenExpired = errors.New("token is expired")
)

func (uc *userUsecaseImpl) Register(ctx context.Context, req *domain.RequestUser) (result domain.ResponseUser, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:Register")
	defer trace.Finish()

	repoFilter := domain.FilterUser{Email: req.Email}
	_, err = uc.repoSQL.UserRepo().Find(ctx, &repoFilter)
	if err == nil {
		err = fmt.Errorf("email %s  already taken", req.Email)
		return
	}

	passHasher := helper.NewPassword(sha1.New, 8, 32, 15000)
	pass := passHasher.HashPassword(req.Password)

	user := shareddomain.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: pass.CipherText,
		PasswordSalt: pass.Salt,
	}

	err = uc.repoSQL.UserRepo().Save(ctx, &user)
	result.Serialize(&user, "")

	// Sample using broker publisher
	// uc.deps.GetBroker(types.Kafka). // get registered broker type (sample Kafka)
	// 				GetPublisher().
	// 				PublishMessage(ctx, &candishared.PublisherArgument{
	// 		Topic:   "[topic]",
	// 		Key:     "[key]",
	// 		Message: candihelper.ToBytes("[message]"),
	// 	})
	return
}

func (uc *userUsecaseImpl) Login(ctx context.Context, req *domain.RequestLoginUser) (result domain.ResponseUser, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:Login")
	defer trace.Finish()

	repoFilter := domain.FilterUser{Email: req.Email}
	user, err := uc.repoSQL.UserRepo().Find(ctx, &repoFilter)
	if err != nil {
		return
	}

	passHasher := helper.NewPassword(sha1.New, 8, 32, 15000)
	if !passHasher.VerifyPassword(req.Password, user.PasswordHash, user.PasswordSalt) {
		err = errors.New("invalid Password")
		return
	}

	payload := authservice.PayloadGenerateToken{DeviceID: "Web", UserID: strconv.Itoa(user.ID)}
	response, _, err := uc.sdk.Authservice().GenerateToken(ctx, payload)
	if err != nil {
		return
	}

	// now := time.Now()
	// exp := now.Add(time.Hour * 1)

	// var key interface{}
	// var token = new(jwt.Token)
	// token = jwt.New(jwt.SigningMethodHS256)
	// key = []byte(TokenKey)

	// claims := jwt.MapClaims{
	// 	"iss": "mooc",
	// 	"exp": exp.Unix(),
	// 	"iat": now.Unix(),
	// 	"did": "Web",
	// 	"aud": "aud",
	// 	"jti": strconv.Itoa(user.ID),
	// 	"sub": strconv.Itoa(user.ID),
	// 	// "user": user,
	// }
	// token.Claims = claims
	// tokenString, err := token.SignedString(key)
	// if err != nil {
	// 	return
	// }

	result.Serialize(&user, response.Token)

	// Sample using broker publisher
	// uc.deps.GetBroker(types.Kafka). // get registered broker type (sample Kafka)
	// 				GetPublisher().
	// 				PublishMessage(ctx, &candishared.PublisherArgument{
	// 		Topic:   "[topic]",
	// 		Key:     "[key]",
	// 		Message: candihelper.ToBytes("[message]"),
	// 	})
	return
}

func (uc *userUsecaseImpl) ValidateToken(ctx context.Context, tokenString string) (claim *candishared.TokenClaim, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:ValidateToken")
	defer trace.Finish()

	resp, _, err := uc.sdk.Authservice().ValidateToken(ctx, tokenString)
	claim = new(candishared.TokenClaim)
	claim.StandardClaims = jwt.StandardClaims{
		Audience:  resp.Audience,
		Subject:   resp.Subject,
		ExpiresAt: resp.ExpiresAt,
		Issuer:    resp.Issuer,
		IssuedAt:  resp.IssuedAt,
		NotBefore: resp.NotBefore,
	}
	claim.Additional = map[string]interface{}{
		"username": resp.User.Username,
		"user_id":  resp.User.ID,
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
