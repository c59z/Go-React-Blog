package config

import (
	"blog-go/model/appTypes"
	"fmt"
	"strings"
)

// System configuration
type System struct {
	Host           string `json:"-" yaml:"host"`                          // Server host, usually 0.0.0.0
	Port           int    `json:"-" yaml:"port"`                          // Server port
	Env            string `json:"-" yaml:"env"`                           // Gin environment: debug, release, or test
	RouterPrefix   string `json:"-" yaml:"router_prefix"`                 // API route prefix
	UseMultipoint  bool   `json:"use_multipoint" yaml:"use_multipoint"`   // Enable multi-point login restriction
	SessionsSecret string `json:"sessions_secret" yaml:"sessions_secret"` // Secret key for session encryption
	OssType        string `json:"oss_type" yaml:"oss_type"`               // Object storage type, e.g., local, qiniu
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s System) Storage() appTypes.Storage {
	switch strings.ToLower(s.OssType) {
	case "local", "Local":
		return appTypes.Local
	case "smms", "SMMS":
		return appTypes.SMMS
	default:
		return appTypes.Local
	}
}
