package usecase

import (
	"context"
	"github.com/google/uuid"
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

func (s StorefrontUseCase) GetProductByUUID(ctx context.Context, productUUID uuid.UUID) (entity.Product, error) {
	//TODO implement me
	panic("implement me")
}
