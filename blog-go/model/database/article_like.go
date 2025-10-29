package database

import "blog-go/global"

// ArticleLike represents article likes/bookmarks
type ArticleLike struct {
	global.MODEL
	ArticleID string `json:"article_id" gorm:"size:191;not null;uniqueIndex:idx_user_article"`
	UserID    uint   `json:"user_id"    gorm:"not null;uniqueIndex:idx_user_article"`
	User      User   `json:"-" gorm:"foreignKey:UserID"`
}
