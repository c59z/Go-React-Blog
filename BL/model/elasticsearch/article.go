package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Article represents the article document in Elasticsearch
type Article struct {
	CreatedAt string `json:"created_at"` // Creation time
	UpdatedAt string `json:"updated_at"` // Last update time

	Cover    string   `json:"cover"`    // Cover image
	Title    string   `json:"title"`    // Title
	Keyword  string   `json:"keyword"`  // Keywords
	Category string   `json:"category"` // Category
	Tags     []string `json:"tags"`     // Tags
	Abstract string   `json:"abstract"` // Abstract
	Content  string   `json:"content"`  // Content

	Views    int `json:"views"`    // View count
	Comments int `json:"comments"` // Comment count
	Likes    int `json:"likes"`    // Like count
}

// ArticleIndex returns the index name for articles
func ArticleIndex() string {
	return "article_index"
}

// ArticleMapping returns the mapping for the article index
func ArticleMapping() *types.TypeMapping {
	return &types.TypeMapping{
		Properties: map[string]types.Property{
			"created_at": types.DateProperty{NullValue: nil, Format: func(s string) *string { return &s }("yyyy-MM-dd HH:mm:ss")},
			"updated_at": types.DateProperty{NullValue: nil, Format: func(s string) *string { return &s }("yyyy-MM-dd HH:mm:ss")},
			"cover":      types.TextProperty{},
			"title":      types.TextProperty{},
			"keyword":    types.KeywordProperty{},
			"category":   types.KeywordProperty{},
			"tags":       []types.KeywordProperty{},
			"abstract":   types.TextProperty{},
			"content":    types.TextProperty{},
			"views":      types.IntegerNumberProperty{},
			"comments":   types.IntegerNumberProperty{},
			"likes":      types.IntegerNumberProperty{},
		},
	}
}
