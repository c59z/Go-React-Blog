package database

// ArticleTag represents article tags
type ArticleTag struct {
	Tag    string `json:"tag" gorm:"primaryKey"` // Tag
	Number int    `json:"number"`                // Count
}
