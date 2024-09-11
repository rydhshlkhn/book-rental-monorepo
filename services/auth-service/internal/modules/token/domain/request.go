package domain

import (
	shareddomain "monorepo/services/auth-service/pkg/shared/domain"
)

// RequestToken model
type RequestToken struct {
	ID    int    `json:"id"`
	Field string `json:"field"`
}

// Deserialize to db model
func (r *RequestToken) Deserialize() (res shareddomain.Token) {
	res.Field = r.Field
	return
}

type RequestGenerateToken struct {
	UserID   string `json:"user_id"`
	DeviceID string `json:"device_id"`
	Role     string `json:"role"`
	Username     string `json:"username"`
}
