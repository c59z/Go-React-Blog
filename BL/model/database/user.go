package database

import (
	"blog-go/global"
	"blog-go/model/appTypes"

	"github.com/gofrs/uuid"
)

// User represents a system user
type User struct {
	global.MODEL
	UUID      uuid.UUID         `json:"uuid" gorm:"type:char(36);unique"` // UUID
	Username  string            `json:"username"`                         // Username
	Password  string            `json:"-"`                                // Password
	Email     string            `json:"email"`                            // Email
	Openid    string            `json:"openid"`                           // OpenID
	Avatar    string            `json:"avatar" gorm:"size:255"`           // Avatar (email/GitHub)
	Address   string            `json:"address"`                          // Address
	Signature string            `json:"signature" gorm:"default:''"`      // Signature
	RoleID    appTypes.RoleID   `json:"role_id"`                          // Role ID
	Register  appTypes.Register `json:"register"`                         // Register source
	Freeze    bool              `json:"freeze"`                           // Frozen status
}
