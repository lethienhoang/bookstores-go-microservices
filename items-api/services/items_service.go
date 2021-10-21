package services

import (
	"encoding/json"
	"github.com/bookstores-go-microservices/items-api/domain"
	"github.com/bookstores-go-microservices/items-api/domain/items"
	"github.com/lethienhoang/bookstores-utils-go/errors"
	"log"
)

var(
	ItemsService IItemServiceInterface = &ItemService{}
)

type IItemServiceInterface interface {
	Create(item *items.Item) (*items.Item, *errors.RestError)
	Query(query interface{}) ([]items.Item, *errors.RestError)
}

type ItemService struct {
}

func (s *ItemService) Create(item *items.Item) (*items.Item, *errors.RestError) {
	if err := item.Save(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	return item, nil
}

func (s *ItemService) Query(query interface{}) ([]items.Item, *errors.RestError) {
	item := items.Item{}
	var results []items.Item
	index := []string{item.GetTableName()}

	res, err := domain.Client.Search(index, query)
	if err != nil {
		log.Fatal(err)
		return nil, errors.NewBadRequestError(err.Error())
	}

	if err:= json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	return results, nil
}