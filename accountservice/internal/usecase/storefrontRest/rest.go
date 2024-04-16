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

func (sr StorefrontRest) GetAllProducts() entity.Products {
	url := sr.cfg.StorefrontUrl
	log.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	var product entity.Products
	jsonDataBytes, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(jsonDataBytes, &product)
	if err != nil {
		log.Println("failed to decode")
		log.Println(err)
	}
	return product
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
