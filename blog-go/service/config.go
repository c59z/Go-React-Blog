package service

import (
	"blog-go/config"
	"blog-go/global"
	"blog-go/model/appTypes"
	"blog-go/utils"

	"gorm.io/gorm"
)

type ConfigService struct {
}

func (configService *ConfigService) UpdateWebsite(website config.Website) error {
	oldArray := []string{
		global.Config.Website.Logo,
		global.Config.Website.FullLogo,
		global.Config.Website.GithubImage,
	}

	newArray := []string{
		website.Logo,
		website.FullLogo,
		website.GithubImage,
	}

	added, removed := utils.DiffArrays(oldArray, newArray)

	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := utils.InitImagesCategory(global.DB, removed); err != nil {
			return err
		}
		if err := utils.ChangeImagesCategory(global.DB, added, appTypes.System); err != nil {
			return err
		}
		global.Config.Website = website
		if err := utils.SaveYAML(); err != nil {
			return err
		}
		return nil
	})
}

func (configService *ConfigService) UpdateSystem(system config.System) error {
	global.Config.System.UseMultipoint = system.UseMultipoint
	global.Config.System.SessionsSecret = system.SessionsSecret
	global.Config.System.OssType = system.OssType
	return utils.SaveYAML()
}

func (configService *ConfigService) UpdateEmail(email config.Email) error {
	global.Config.Email = email
	return utils.SaveYAML()
}

func (configService *ConfigService) UpdateGithub(github config.GitHub) error {
	global.Config.Github = github
	return utils.SaveYAML()
}

func (configService *ConfigService) UpdateJwt(jwt config.Jwt) error {
	global.Config.Jwt = jwt
	return utils.SaveYAML()
}
