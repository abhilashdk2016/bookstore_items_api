package services

import (
	"github.com/abhilashdk2016/bookstore_items_api/domain/items"
	"github.com/abhilashdk2016/bookstore_items_api/domain/queries"
	"github.com/abhilashdk2016/bookstore_utils_go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService { }
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemsService struct {}

func (s *itemsService) Create(itemRequest items.Item) (*items.Item, rest_errors.RestErr) {
	if err := itemRequest.Save(); err != nil {
		return nil, err
	}
	return &itemRequest, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
