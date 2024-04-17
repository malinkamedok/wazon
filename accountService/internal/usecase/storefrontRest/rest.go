package storefrontrest

import (
	"accountservice/internal/config"
	"accountservice/internal/entity"
	"encoding/json"
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
	log.Println(url)

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

//func GetUniqueProduct() entity.Products {
//	url := "http://localhost:8082/storefront/"
//	req, _ := http.NewRequest("GET", url, nil)
//	req.Header.Add("accept", "application/json")
//	res, _ := http.DefaultClient.Do(req)
//	defer res.Body.Close()
//	var product entity.Products
//	jsonDataBytes, _ := io.ReadAll(res.Body)
//	err := json.Unmarshal(jsonDataBytes, &product)
//	if err != nil {
//		log.Println("failed to decode")
//		log.Println(err)
//	}
//	return product
//}
