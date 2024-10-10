package domain

const (
	// RedisTokenExpiredKeyConst const
	RedisTokenExpiredKeyConst = "expiredtoken"
)

type ResponseGenerateToken struct {
	Token        string                 `json:"token"`
	RefreshToken string                 `json:"refresh_token"`
	Claim        map[string]interface{} `json:"claim"`
}
