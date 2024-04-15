package integration

import (
	"accountservice/internal/entity"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetAllProducts() entity.Products {
	url := "http://localhost:8082/storefront/" //TODO как это достать из конфига?
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
//	url := "http://localhost:8082/storefront/" //TODO как это достать из конфига?
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
