package config

type Captcha struct {
	Height   int     `json:"height" yaml:"height"`       // Height of the PNG image in pixels
	Width    int     `json:"width" yaml:"width"`         // Width of the verification code PNG image in pixels
	Length   int     `json:"length" yaml:"length"`       // The default number of digits in the verification code result
	MaxSkew  float64 `json:"max_skew" yaml:"max_skew"`   // Maximum skew factor of a single number (absolute value)
	DotCount int     `json:"dot_count" yaml:"dot_count"` // Number of background dots
}
