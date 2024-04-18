package usecase

import (
	"accountservice/internal/entity"
	"context"
	"github.com/google/uuid"
)

type (
	AccountServiceContract interface {
		GetUserById(ctx context.Context, userId uuid.UUID) (entity.User, error)
		InsertOrUpdateProduct(ctx context.Context, product entity.Product) error
		GetAllProductsFromCart(ctx context.Context, userId uuid.UUID) ([]entity.Product, error)
		CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error)
	}

	AccountServiceRepository interface {
		GetUserById(ctx context.Context, userId uuid.UUID) (entity.User, error)
		InsertOrUpdateProduct(ctx context.Context, product entity.Product) error
		GetAllProductsFromCart(ctx context.Context, userId uuid.UUID) ([]entity.Product, error)
		CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error)
	}

	IntegrationRest interface {
		GetAllProducts() (entity.Products, error)
	}

	IntegrationContract interface {
		GetAllProducts() (entity.Products, error)
	}
)
