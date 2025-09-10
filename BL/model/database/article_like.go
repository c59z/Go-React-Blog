package database

import "blog-go/global"

// ArticleLike represents article likes/bookmarks
type ArticleLike struct {
	global.MODEL
	ArticleID string `json:"article_id"` // Article ID
	UserID    uint   `json:"user_id"`    // User ID
	User      User   `json:"-" gorm:"foreignKey:UserID"`
}
