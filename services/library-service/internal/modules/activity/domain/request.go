package domain

import (
	shareddomain "monorepo/services/library-service/pkg/shared/domain"
)

// RequestActivity model
type RequestActivity struct {
	ID    int `json:"id"`
	Field string `json:"field"`
}

// Deserialize to db model
func (r *RequestActivity) Deserialize() (res shareddomain.Activity) {
	res.Field = r.Field
	return
}
