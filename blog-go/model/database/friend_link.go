package database

import "blog-go/global"

// FriendLink table
type FriendLink struct {
	global.MODEL
	Logo        string `json:"logo" gorm:"size:255"`                    // Logo URL
	Image       Image  `json:"-" gorm:"foreignKey:Logo;references:URL"` // Related image
	Link        string `json:"link"`                                    // Website link
	Name        string `json:"name"`                                    // Site name
	Description string `json:"description"`                             // Site description
}
