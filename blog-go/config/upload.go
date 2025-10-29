package config

type Upload struct {
	Size int    `json:"size" yaml:"size"` // Max upload size (MB)
	Path string `json:"path" yaml:"path"` // Upload directory
}
