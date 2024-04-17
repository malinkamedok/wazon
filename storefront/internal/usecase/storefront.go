package usecase

import (
	"context"
	"github.com/google/uuid"
	"log"
	"storefront/internal/entity"
)

type StorefrontUseCase struct {
	repo StorefrontRepository
}

var _ StorefrontContract = (*StorefrontUseCase)(nil)

func NewStorefrontUseCase(repo StorefrontRepository) *StorefrontUseCase {
	return &StorefrontUseCase{repo: repo}
}

func (s StorefrontUseCase) GetAllProducts(ctx context.Context) ([]entity.ProductList, error) {
	products, err := s.repo.ReadAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s StorefrontUseCase) GetProductByUUID(ctx context.Context, productIDStr string) (entity.Product, error) {
	productUUID, err := uuid.Parse(productIDStr)
	if err != nil {
		log.Println("could not parse product uuid")
		return entity.Product{}, err
	}
	product, err := s.repo.ReadProductByUUID(ctx, productUUID)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}
