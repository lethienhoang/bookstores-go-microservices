package services

import (
	"fmt"
	"github.com/bookstores-go-microservices/items-api/domain/items"
	"github.com/lethienhoang/bookstores-utils-go/errors"
)

var(
	ItemsService IItemServiceInterface = &ItemService{}
)

type IItemServiceInterface interface {
	Create(item *items.Item) (*items.Item, *errors.RestError)
	Get(string) (*items.Item, *errors.RestError)
}

type ItemService struct {
}

func (s *ItemService) Create(item *items.Item) (*items.Item, *errors.RestError) {
	fmt.Println(item)
	if err := item.Save(); err != nil {
		return nil, errors.NewBadRequestError(err.Error())
	}

	return item, nil

}

func (s *ItemService) Get(string) (*items.Item, *errors.RestError) {
	return nil, nil
}