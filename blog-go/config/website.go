package config

// Website info
type Website struct {
	Logo        string `json:"logo" yaml:"logo"`
	FullLogo    string `json:"full_logo" yaml:"full_logo"`
	Title       string `json:"title" yaml:"title"`               // Website title
	Slogan      string `json:"slogan" yaml:"slogan"`             // Website slogan
	SloganEn    string `json:"slogan_en" yaml:"slogan_en"`       // English slogan
	Description string `json:"description" yaml:"description"`   // Website description
	Version     string `json:"version" yaml:"version"`           // Website version
	CreatedAt   string `json:"created_at" yaml:"created_at"`     // Creation date
	BilibiliURL string `json:"bilibili_url" yaml:"bilibili_url"` // Bilibili link
	GithubURL   string `json:"github_url" yaml:"github_url"`     // GitHub link
	Name        string `json:"name" yaml:"name"`                 // Author name/nickname
	Job         string `json:"job" yaml:"job"`                   // Job title
	Address     string `json:"address" yaml:"address"`           // Address
	Email       string `json:"email" yaml:"email"`               // Contact email
	GithubImage string `json:"github_image" yaml:"github_image"` // Github Image
}
