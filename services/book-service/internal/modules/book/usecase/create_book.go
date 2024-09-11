package usecase

import (
	"context"

	"monorepo/services/book-service/internal/modules/book/domain"

	"github.com/golang-jwt/jwt"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *bookUsecaseImpl) CreateBook(ctx context.Context, req *domain.RequestBook) (result domain.ResponseBook, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:CreateBook")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.BookRepo().SaveBook(ctx, &data)
	result.Serialize(&data)

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

func (uc *bookUsecaseImpl) ValidateToken(ctx context.Context, tokenString string) (claim *candishared.TokenClaim, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookUsecase:ValidateToken")
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