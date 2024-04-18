package storefrontrest

import (
	"accountservice/internal/config"
	"accountservice/internal/entity"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
)

type StorefrontRest struct {
	cfg *config.Config
}

func NewStorefrontRest(cfg *config.Config) *StorefrontRest {
	return &StorefrontRest{cfg}
}

func (sr StorefrontRest) GetAllProducts() (entity.Products, error) {
	url := sr.cfg.StorefrontUrl + sr.cfg.StoreFrontPort + "/storefront"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("failed to create http request")
		return entity.Products{}, err
	}

	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("failed to do http request")
		return entity.Products{}, err
	}
	defer res.Body.Close()

	var product entity.Products
	jsonDataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("failed to parse http response")
		return entity.Products{}, err
	}

	err = json.Unmarshal(jsonDataBytes, &product)
	if err != nil {
		log.Println("failed to decode")
		log.Println(err)
		return entity.Products{}, err
	}
	return product, nil
}

type orderRequest struct {
	UuidOrder uuid.UUID `json:"uuid"`
}

type orderResponse struct {
	Order   entity.Order `json:"order"`
	Service string       `json:"service"`
}

func (sr StorefrontRest) CreateOrder(uuid uuid.UUID) (entity.Order, error) {
	url := sr.cfg.DeliveryUrl + sr.cfg.DeliveryPort + "/delivery/create"
	body, err := json.Marshal(orderRequest{UuidOrder: uuid})
	if err != nil {
		log.Println("failed to marshal body")
		return entity.Order{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("failed to create http request")
		return entity.Order{}, err
	}

	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("failed to do http request")
		return entity.Order{}, err
	}

	var order orderResponse
	jsonDataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("failed to parse http response")
		return entity.Order{}, err
	}

	err = json.Unmarshal(jsonDataBytes, &order)
	if err != nil {
		log.Println("failed to decode")
		log.Println(err)
	}

	return order.Order, nil
}
