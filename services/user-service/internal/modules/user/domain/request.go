package domain

import (
	shareddomain "monorepo/services/user-service/pkg/shared/domain"
)

// RequestUser model
type RequestUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Deserialize to db model
func (r *RequestUser) Deserialize() (res shareddomain.User) {
	res.Username = r.Username
	return
}

// RequestUser model
type RequestLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
