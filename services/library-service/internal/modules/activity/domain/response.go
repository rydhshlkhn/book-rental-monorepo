package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseActivityList model
type ResponseActivityList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseActivity   `json:"data"`
}

// ResponseActivity model
type ResponseActivity struct {
	ID        int `json:"id"`
	Field     string `json:"field"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseActivity) Serialize(source *shareddomain.Activity) {
	r.ID = source.ID
	r.Field = source.Field
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
