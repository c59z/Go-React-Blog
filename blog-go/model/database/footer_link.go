package database

// FooterLink table
type FooterLink struct {
	Title string `json:"title" gorm:"primaryKey"` // Link title
	Link  string `json:"link"`                    // Link URL
}
