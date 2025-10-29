package service

import (
	"blog-go/global"
	"blog-go/model/database"
	"blog-go/utils"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type JwtService struct {
}

func (j *JwtService) SetRedisJwt(jwt string, uuid uuid.UUID) error {
	dr, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	if err != nil {
		return err
	}
	return global.Redis.Set(uuid.String(), jwt, dr).Err()
}

func (j *JwtService) GetRedisJwt(uuid uuid.UUID) (string, error) {
	return global.Redis.Get(uuid.String()).Result()
}

func (j *JwtService) JoinInBlacklist(jwtList database.JwtBlacklist) error {
	if err := global.DB.Create(&jwtList).Error; err != nil {
		return err
	}

	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return nil
}

func (j *JwtService) IsInBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

func LoadAllFromJwtBlackList() {
	var data []string

	if err := global.DB.Model(&database.JwtBlacklist{}).Pluck("jwt", &data).Error; err != nil {
		global.Log.Error("Failed to load JWT blacklist from the database", zap.Error(err))
		return
	}

	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
