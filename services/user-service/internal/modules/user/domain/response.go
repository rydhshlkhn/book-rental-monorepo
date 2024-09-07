package domain

import (
	shareddomain "monorepo/services/user-service/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseUserList model
type ResponseUserList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseUser   `json:"data"`
}

// ResponseUser model
type ResponseUser struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseUser) Serialize(source *shareddomain.User, token string) {
	r.ID = source.ID
	r.Username = source.Username
	r.Email = source.Email
	r.Token = token
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}
