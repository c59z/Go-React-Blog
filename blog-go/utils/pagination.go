package utils

import (
	"blog-go/global"
	"blog-go/model/other"
	"context"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// MySQLPagination implements pagination for MySQL queries using GORM
func MySQLPagination[T any](model *T, option other.MySQLOption) (list []T, total int64, err error) {
	// Set default pagination values
	if option.Page < 1 {
		option.Page = 1 // Page number cannot be less than 1, default to 1
	}
	if option.PageSize < 1 {
		option.PageSize = 10 // Page size cannot be less than 1, default to 10
	}
	if option.Order == "" {
		option.Order = "id desc" // Default order by ID in descending order
	}

	// Create the initial query
	query := global.DB.Model(model)

	// Apply additional WHERE conditions if provided
	if option.Where != nil {
		query = query.Where(option.Where)
	}

	// Count total records that match the conditions
	if err = query.Count(&total).Error; err != nil {
		return nil, 0, err // Return error if counting fails
	}

	// Preload related models if specified
	for _, preload := range option.Preload {
		query = query.Preload(preload) // Apply eager loading for related associations
	}

	// Apply pagination and execute the query
	err = query.Order(option.Order).
		Limit(option.PageSize).                      // Set number of records per page
		Offset((option.Page - 1) * option.PageSize). // Calculate offset based on page number
		Find(&list).Error                            // Execute query and store results in list

	return list, total, err // Return paginated results and total count
}

func EsPagination(ctx context.Context, option other.EsOption) (list []types.Hit, total int64, err error) {
	// Set default pagination values
	if option.Page < 1 {
		option.Page = 1 // Page number must be at least 1
	}
	if option.PageSize < 1 {
		option.PageSize = 10 // Default page size is 10
	}

	// Set pagination for Elasticsearch query
	from := (option.Page - 1) * option.PageSize // Calculate start position
	option.Request.Size = &option.PageSize      // Set page size
	option.Request.From = &from                 // Set start record position

	// Execute Elasticsearch search
	res, err := global.ESClient.Search().
		Index(option.Index).                       // Specify index
		Request(option.Request).                   // Apply query request
		SourceIncludes_(option.SourceIncludes...). // Include selected fields
		Do(ctx)                                    // Execute query
	if err != nil {
		return nil, 0, err // Return error if search fails
	}

	// Get search results
	list = res.Hits.Hits         // List of matched documents
	total = res.Hits.Total.Value // Total number of matched documents
	return list, total, nil      // Return list and total count
}
