package usecase

import (
	"accountservice/internal/entity"
	storefrontrest "accountservice/internal/usecase/storefrontRest"
)

type IntegrationUsecase struct {
	rest *storefrontrest.StorefrontRest
}

func NewIntegrationUsecase(rest *storefrontrest.StorefrontRest) *IntegrationUsecase {
	return &IntegrationUsecase{rest: rest}
}

func (i IntegrationUsecase) GetAllProducts() entity.Products {
	return i.rest.GetAllProducts()
}

var _ IntegrationContract = (*IntegrationUsecase)(nil)
