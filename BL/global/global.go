package global

import (
	"blog-go/config"

	"go.uber.org/zap"
)

var (
	Config *config.Config
	Log    *zap.Logger
)
