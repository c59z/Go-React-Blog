package utils

import (
	"blog-go/model/appTypes"
	"blog-go/model/database"

	"gorm.io/gorm"
)

func InitImagesCategory(tx *gorm.DB, urls []string) error {
	return tx.Model(&database.Image{}).Where("url IN ?", urls).Update("category", appTypes.Null).Error
}

func ChangeImagesCategory(tx *gorm.DB, urls []string, category appTypes.Category) error {
	return tx.Model(&database.Image{}).Where("url IN ?", urls).Update("category", category).Error
}
