package upload

import (
	"blog-go/global"
	"blog-go/model/appTypes"
	"mime/multipart"
)

var WhiteImageList = map[string]struct{}{
	".jpg":  {},
	".png":  {},
	".jpeg": {},
	".ico":  {},
	".tiff": {},
	".gif":  {},
	".svg":  {},
	".webp": {},
}

type OSS interface {
	UploadImage(file *multipart.FileHeader) (string, string, error)
	DeleteImage(key string) error
}

func NewOss() OSS {
	switch global.Config.System.OssType {
	case "local":
		return &Local{}
	case "smms":
		return &Local{}
	default:
		return &Local{}
	}
}

func NewOssWithStorage(storage appTypes.Storage) OSS {
	switch storage {
	case appTypes.Local:
		return &Local{}
	case appTypes.SMMS:
		return &Local{}
	default:
		return &Local{}
	}
}
