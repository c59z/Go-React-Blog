package database

import (
	"blog-go/global"
	"blog-go/model/appTypes"
)

// Image represents stored images
type Image struct {
	global.MODEL
	Name     string            `json:"name"`                       // Name
	URL      string            `json:"url" gorm:"size:255;unique"` // URL
	Category appTypes.Category `json:"category"`                   // Category
	Storage  appTypes.Storage  `json:"storage"`                    // Storage type
}
