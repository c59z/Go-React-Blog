package flag

import (
	"blog-go/global"
	"blog-go/model/elasticsearch"
	"blog-go/model/other"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ElasticsearchExport exports data from Elasticsearch to a JSON file
func ElasticsearchExport() error {
	// Response container
	var response other.ESIndexResponse

	// Initial search request: index name, scroll time, batch size, query (match all)
	res, err := global.ESClient.Search().
		Index(elasticsearch.ArticleIndex()).                   // Index name
		Scroll("1m").                                          // Scroll time: 1 minute
		Size(1000).                                            // Number of docs per batch
		Query(&types.Query{MatchAll: &types.MatchAllQuery{}}). // Match all documents
		Do(context.TODO())                                     // Execute with empty context
	if err != nil {
		return err
	}

	// Process initial search results
	for _, hit := range res.Hits.Hits {
		data := other.Data{
			ID:  hit.Id_,
			Doc: hit.Source_,
		}
		response.Data = append(response.Data, data)
	}

	// Continue scrolling until no more results
	for {
		res, err := global.ESClient.Scroll().ScrollId(*res.ScrollId_).Scroll("1m").Do(context.TODO())
		if err != nil {
			return err
		}

		// Break if no more documents
		if len(res.Hits.Hits) == 0 {
			break
		}

		// Process documents in this batch
		for _, hit := range res.Hits.Hits {
			data := other.Data{
				ID:  hit.Id_,
				Doc: hit.Source_,
			}
			response.Data = append(response.Data, data)
		}
	}

	// Clear scroll to free resources
	_, err = global.ESClient.ClearScroll().ScrollId(*res.ScrollId_).Do(context.TODO())
	if err != nil {
		return err
	}

	// Generate file name: "es_yyyyMMdd.json"
	fileName := fmt.Sprintf("es_%s.json", time.Now().Format("20060102"))

	// Create output file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Marshal response to JSON
	byteData, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// Write JSON to file
	_, err = file.Write(byteData)
	if err != nil {
		return err
	}

	return nil
}
