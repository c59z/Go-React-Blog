package database

import (
	"blog-go/global"

	"github.com/gofrs/uuid"
)

// Feedback table
type Feedback struct {
	global.MODEL
	UserUUID uuid.UUID `json:"user_uuid" gorm:"type:char(36)"`               // User UUID
	User     User      `json:"-" gorm:"foreignKey:UserUUID;references:UUID"` // Related user
	Content  string    `json:"content"`                                      // Feedback content
	Reply    string    `json:"reply"`                                        // Admin reply
}
