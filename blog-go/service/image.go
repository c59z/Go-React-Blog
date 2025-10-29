package service

import (
	"blog-go/global"
	"blog-go/model/appTypes"
	"blog-go/model/database"
	"blog-go/model/other"
	"blog-go/model/request"
	"blog-go/utils"
	"blog-go/utils/upload"
	"mime/multipart"

	"gorm.io/gorm"
)

type ImageService struct {
}

func (s *ImageService) ImageUpload(file *multipart.FileHeader) (string, error) {
	oss := upload.NewOss()
	url, filename, err := oss.UploadImage(file)
	if err != nil {
		return "", err
	}
	return url, global.DB.Create(&database.Image{
		Name:     filename,
		URL:      url,
		Category: appTypes.Null,
		Storage:  global.Config.System.Storage(),
	}).Error
}
func (s *ImageService) ImageDelete(req request.ImageDelete) error {
	if len(req.IDs) == 0 {
		return nil
	}

	var images []database.Image
	if err := global.DB.Find(&images, req.IDs).Error; err != nil {
		return err
	}
	for _, image := range images {
		if err := global.DB.Transaction(func(tx *gorm.DB) error {
			oss := upload.NewOssWithStorage(image.Storage)
			if err := tx.Delete(&image).Error; err != nil {
				return err
			}
			return oss.DeleteImage(image.Name)
		}); err != nil {
			return err
		}
	}
	return nil
}
func (s *ImageService) ImageList(req request.ImageList) (interface{}, int64, error) {
	db := global.DB

	if req.Name != nil {
		db = db.Where("name LIKE ?", "%"+*req.Name+"%")
	}

	if req.Category != nil {
		category := appTypes.ToCategory(*req.Category)
		db = db.Where("category = ?", category)
	}

	if req.Storage != nil {
		storage := appTypes.ToStorage(*req.Storage)
		db = db.Where("storage = ?", storage)
	}

	option := other.MySQLOption{
		PageInfo: req.PageInfo,
		Where:    db,
	}
	return utils.MySQLPagination(&database.Image{}, option)
}
