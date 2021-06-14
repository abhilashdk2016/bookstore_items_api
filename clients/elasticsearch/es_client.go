package elasticsearch

import (
	"context"
	"fmt"
	"github.com/abhilashdk2016/bookstore_utils_go/logger"
	"gopkg.in/olivere/elastic.v7"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c * elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string) (*elastic.GetResult, error)
	Search(index string, query elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10 * time.Second),
		//elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)

	if err != nil {
		panic(err)
	}
    fmt.Println("Elastic Search Connected!!!")
	Client.setClient(client)
}

func (e *esClient) setClient(client * elastic.Client) {
	e.client = client
}

func (e *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := e.client.Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error while trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (e *esClient) Get(index string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := e.client.Get().Index(index).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error while trying to get document %s in index %s", id, index), err)
		return nil, err
	}

	return result, nil
}

func (e *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := e.client.Search(index).Query(query).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search document with index : %s", index), err)
		return nil, err
	}
	return result, nil
}