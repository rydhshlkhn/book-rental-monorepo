package domain

import (
	shareddomain "monorepo/services/auth-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseTokenList model
type ResponseTokenList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseToken   `json:"data"`
}

// ResponseToken model
type ResponseToken struct {
	ID        int `json:"id"`
	Field     string `json:"field"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseToken) Serialize(source *shareddomain.Token) {
	r.ID = source.ID
	r.Field = source.Field
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
