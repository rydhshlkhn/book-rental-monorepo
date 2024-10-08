// Code generated by candi v1.17.15.

package graphqlhandler

import (
	"context"

	"monorepo/services/library-service/internal/modules/book/domain"

	"github.com/golangid/candi/tracer"
)

// CreateBook resolver
func (m *GraphQLHandler) CreateBook(ctx context.Context, input struct{ Data domain.RequestBook }) (data domain.ResponseBook, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookDeliveryGraphQL:CreateBook")
	defer trace.Finish()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	if err := m.validator.ValidateDocument("book/save", input.Data); err != nil {
		return data, err
	}
	return m.uc.Book().CreateBook(ctx, &input.Data)
}

// UpdateBook resolver
func (m *GraphQLHandler) UpdateBook(ctx context.Context, input struct {
	ID   int
	Data domain.RequestBook
}) (ok string, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookDeliveryGraphQL:UpdateBook")
	defer trace.Finish()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	input.Data.ID = input.ID
	if err := m.validator.ValidateDocument("book/save", input.Data); err != nil {
		return "", err
	}
	if err := m.uc.Book().UpdateBook(ctx, &input.Data); err != nil {
		return ok, err
	}
	return "Success", nil
}

// DeleteBook resolver
func (m *GraphQLHandler) DeleteBook(ctx context.Context, input struct{ ID int }) (ok string, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BookDeliveryGraphQL:DeleteBook")
	defer trace.Finish()

	// tokenClaim := candishared.ParseTokenClaimFromContext(ctx) // must using GraphQLBearerAuth in middleware for this resolver

	if err := m.uc.Book().DeleteBook(ctx, input.ID); err != nil {
		return ok, err
	}
	return "Success", nil
}
