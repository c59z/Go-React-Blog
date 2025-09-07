package config

// Email configuration. For example, enable third-party services in QQ Mail.
// See https://mail.qq.com/
type Email struct {
	Host     string `json:"host" yaml:"host"`         // Mail server address, e.g., smtp.qq.com
	Port     int    `json:"port" yaml:"port"`         // Mail server port, e.g., 587 (TLS) or 465 (SSL)
	From     string `json:"from" yaml:"from"`         // Sender email address
	Nickname string `json:"nickname" yaml:"nickname"` // Sender display name
	Secret   string `json:"secret" yaml:"secret"`     // Password or app-specific password
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl"`     // Use SSL encryption if true
}
