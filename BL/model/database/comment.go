package database

import (
	"blog-go/global"

	"github.com/gofrs/uuid"
)

type Comment struct {
	global.MODEL
	ArticleID string    `json:"article_id"`
	PID       *uint     `json:"p_id"`
	PComment  *Comment  `json:"-" gorm:"foreignKey:PID"`
	Children  []Comment `json:"children" gorm:"foreignKey:PID"`
	UserUUID  uuid.UUID `json:"user_uuid" gorm:"type:char(36)"`
	User      User      `json:"user" gorm:"foreignKey:UserUUID;references:UUID"`
	Content   string    `json:"content"`
}
