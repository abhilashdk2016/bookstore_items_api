package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abhilashdk2016/bookstore_items_api/clients/elasticsearch"
	"github.com/abhilashdk2016/bookstore_items_api/domain/queries"
	"github.com/abhilashdk2016/bookstore_utils_go/rest_errors"
	"strings"
)

const (
	itemsIndex = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(itemsIndex, i)
	if err != nil {
		return rest_errors.NewInternalServerError("Error while trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	result, err := elasticsearch.Client.Get(itemsIndex, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("No item found with id: %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("Error while trying to get id %s", i.Id), errors.New("Database Error"))
	}

	if !result.Found {
		return rest_errors.NewNotFoundError(fmt.Sprintf("No item found with id: %s", i.Id))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("Error while trying to parse database response", errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerError("Error while trying to parse database response", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(itemsIndex, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error while trying to search documents", errors.New("es error"))
	}
	// fmt.Println(result)
	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error while trying to parse response", errors.New("Database error"))
		}
		// item.Id = hit.Id
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}
