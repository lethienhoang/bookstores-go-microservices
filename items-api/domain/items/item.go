package items

import (
	"github.com/bookstores-go-microservices/items-api/domain"
	"log"
)

type Item struct {
	Id               string      `json:"id"`
	Seller           int64       `json:"seller"`
	Title            string      `json:"title"`
	Description      Description `json:"description"`
	Pictures         []Picture   `json:"pictures"`
	Video            string      `json:"video"`
	Price            float32     `json:"price"`
	AvaiableQuantity int         `json:"avaiable_quantity"`
	SoldQuantity     int         `json:"sold_quantity"`
	Status           string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}

func (t *Item) GetTableName() string {
	return "items"
}

func (t *Item) Save() error  {
	_, err := domain.Client.Index(t.GetTableName(), t.Id, t)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
