package config

// Google Maps service configuration
type Google struct {
	Enable bool   `json:"enable" yaml:"enable"` // Whether to enable Google Maps service
	Key    string `json:"key" yaml:"key"`       // Google Maps API key
}
