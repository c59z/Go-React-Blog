package flag

import (
	"blog-go/global"
	"blog-go/model/elasticsearch"
	"blog-go/model/other"
	"blog-go/service"
	"context"
	"encoding/json"
	"os"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

// ElasticsearchImport imports data from a JSON file into Elasticsearch
func ElasticsearchImport(jsonPath string) (int, error) {
	// Read JSON file from the given path
	byteData, err := os.ReadFile(jsonPath)
	if err != nil {
		return 0, err
	}

	// Unmarshal JSON data into ESIndexResponse struct
	var response other.ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		return 0, err
	}

	// Create Elasticsearch index
	esService := service.ServiceGroupApp.EsService
	indexExists, err := esService.IndexExists(elasticsearch.ArticleIndex())
	if err != nil {
		return 0, err
	}
	if indexExists {
		if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
			return 0, err
		}
	}
	err = esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping())
	if err != nil {
		return 0, err
	}

	// Build bulk request data
	var request bulk.Request
	for _, data := range response.Data {
		// Add index operation with document ID
		request = append(request, types.OperationContainer{Index: &types.IndexOperation{Id_: data.ID}})
		// Add document data
		request = append(request, data.Doc)
	}

	// Execute bulk operation using Elasticsearch client
	_, err = global.ESClient.Bulk().
		Request(&request).                   // Set request data
		Index(elasticsearch.ArticleIndex()). // Set index name
		Refresh(refresh.True).               // Force refresh to make documents searchable immediately
		Do(context.TODO())                   // Execute request
	if err != nil {
		return 0, err
	}

	// Return total number of imported documents
	total := len(response.Data)
	return total, nil
}
