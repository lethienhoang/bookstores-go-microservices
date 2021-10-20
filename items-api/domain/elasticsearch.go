package domain

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

var (
	Client elasticsearchInterface = &elasticsearchClient{}
)

type elasticsearchInterface interface {
	Index(index string, id string, body interface{}) (*esapi.Response, error)

}

type elasticsearchClient struct {
	client *elasticsearch.Client
}

func NewElasticSearchClient() *elasticsearchClient {
	cfg := elasticsearch.Config{
		Addresses: []string {
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	return &elasticsearchClient {
		client: es,
	}
}

func (c *elasticsearchClient) Index(index string, id string, body interface{}) (*esapi.Response, error)  {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Error converting structure to json: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(string(bodyJson)),
		Refresh:    "true",
	}

	ctx := context.Background()
	res, err := req.Do(ctx, c.client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	return res, nil
}

func (c *elasticsearchClient) Search (index []string, query interface{}) (*esapi.Response, error) {
	bodyJson, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error converting structure to json: %s", err)
	}

	req := esapi.SearchRequest{
		Index:      index,
		Body:       strings.NewReader(string(bodyJson)),
		TrackTotalHits:    true,
		Pretty: true,
	}

	ctx := context.Background()
	res, err := req.Do(ctx, c.client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	return res, nil

}