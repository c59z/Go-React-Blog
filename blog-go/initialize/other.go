package initialize

import (
	"blog-go/global"
	"blog-go/utils"
	"os"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
)

func OtherInit() {
	// Parse refresh token expiry time
	refreshTokenExpiry, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse refresh token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	// Parse access token expiry time
	_, err = utils.ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)
	if err != nil {
		global.Log.Error("Failed to parse access token expiry time configuration:", zap.Error(err))
		os.Exit(1)
	}

	// Configure local cache expiration (using refresh token expiry time
	// to facilitate JWT blacklist handling in cases such as remote login
	// or account suspension)
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(refreshTokenExpiry),
	)
}
