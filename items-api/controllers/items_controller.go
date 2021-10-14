package controllers

import (
	"encoding/json"
	"github.com/bookstores-go-microservices/items-api/domain/items"
	"github.com/bookstores-go-microservices/items-api/services"
	"github.com/lethienhoang/bookstores-utils-go/jwt_auth"
	"io/ioutil"
	"net/http"
)

var (
	ItemsController IItemsControllerInterface = &ItemController{}
)

type IItemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	//Search(w http.ResponseWriter, r *http.Request)
}

type ItemController struct {
}

func (s *ItemController) Create(w http.ResponseWriter, r *http.Request) {
	bearToken := r.Header.Get("Authorization")
	decodeToken, errToken := jwt_auth.DecodeToken(bearToken, false)
	if errToken != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errToken.Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	defer r.Body.Close()

	var itemReq items.Item
	if err := json.Unmarshal(body, &itemReq); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	itemReq.Seller = decodeToken.UserId

	result, restErr := services.ItemsService.Create(&itemReq)
	if restErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(restErr.Code)
		json.NewEncoder(w).Encode(restErr.Message)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (s *ItemController) Get(w http.ResponseWriter, r *http.Request) {

}
