package usecase

import (
	"accountservice/internal/entity"
	storefrontrest "accountservice/internal/usecase/storefrontRest"
	"github.com/google/uuid"
	"log"
)

type IntegrationUsecase struct {
	rest *storefrontrest.StorefrontRest
}

func (i IntegrationUsecase) CreateOrder(uuidStr string) (entity.Order, error) {
	uuidConverted, err := uuid.Parse(uuidStr)
	if err != nil {
		log.Println("error in parsing uuid ", err)
		return entity.Order{}, err
	}
	return i.rest.CreateOrder(uuidConverted)
}

func NewIntegrationUsecase(rest *storefrontrest.StorefrontRest) *IntegrationUsecase {
	return &IntegrationUsecase{rest: rest}
}

func (i IntegrationUsecase) GetAllProducts() (entity.Products, error) {
	return i.rest.GetAllProducts()
}

var _ IntegrationContract = (*IntegrationUsecase)(nil)
