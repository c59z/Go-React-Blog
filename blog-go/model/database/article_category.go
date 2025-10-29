package database

// ArticleCategory represents article categories
type ArticleCategory struct {
	Category string `json:"category" gorm:"primaryKey"` // Category name
	Number   int    `json:"number"`                     // Count
}
