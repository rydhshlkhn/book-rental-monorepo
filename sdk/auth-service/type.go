package authservice

import "github.com/golang-jwt/jwt"

type PayloadGenerateToken struct {
	DeviceID string
	UserID   string
	Username string
	Role string
}

type ResponseGenerateToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseClaim struct {
	jwt.StandardClaims
	DeviceID string `json:"did"`
	User     struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
	Alg string `json:"-"`
}
