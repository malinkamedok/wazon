package usecase

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"storefront/internal/entity"
	"storefront/pkg/logger"
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
		logger.Error("could not parse product uuid", zap.Error(err))
		return entity.Product{}, err
	}
	product, err := s.repo.ReadProductByUUID(ctx, productUUID)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}
