package config

// Jwt configuration
type Jwt struct {
	AccessTokenSecret      string `json:"access_token_secret" yaml:"access_token_secret"`             // Secret key used to generate and verify access tokens
	RefreshTokenSecret     string `json:"refresh_token_secret" yaml:"refresh_token_secret"`           // Secret key used to generate and verify refresh tokens
	AccessTokenExpiryTime  string `json:"access_token_expiry_time" yaml:"access_token_expiry_time"`   // Expiry time for access tokens, e.g., "15m" means 15 minutes
	RefreshTokenExpiryTime string `json:"refresh_token_expiry_time" yaml:"refresh_token_expiry_time"` // Expiry time for refresh tokens, e.g., "30d" means 30 days
	Issuer                 string `json:"issuer" yaml:"issuer"`                                       // JWT issuer information, usually the application or service name
}
