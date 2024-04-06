package usecase

import (
	"context"
	"github.com/google/uuid"
	"storefront/internal/entity"
)

type (
	StorefrontContract interface {
		GetAllProducts(ctx context.Context) ([]entity.ProductList, error)
		GetProductByUUID(ctx context.Context, productUUID uuid.UUID) (entity.Product, error)
	}

	StorefrontRepository interface {
		ReadAllProducts(ctx context.Context) ([]entity.ProductList, error)
		ReadProductByUUID(ctx context.Context, productUUID uuid.UUID) (entity.Product, error)
	}
)
