package domain

import "github.com/golang-jwt/jwt"

const (

	// HS256 const
	HS256 = "HS256"

	// RS256 const
	RS256 = "RS256"
)

type Claim struct {
	jwt.StandardClaims
	DeviceID string `json:"did"`
	User     struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Role     string `json:"role"`
	} `json:"user"`
	Alg string `json:"-"`
}
