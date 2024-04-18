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
		AddProductToCart(ctx context.Context, userId string, productId string) error
	}

	AccountServiceRepository interface {
		GetUserById(ctx context.Context, userId uuid.UUID) (entity.User, error)
		InsertOrUpdateProduct(ctx context.Context, product entity.Product) error
		GetAllProductsFromCart(ctx context.Context, userId uuid.UUID) ([]entity.Product, error)
		CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error)

		AddProductToCart(ctx context.Context, cartID uuid.UUID, productId uuid.UUID) error
		CreateCart(ctx context.Context, userId uuid.UUID) (uuid.UUID, error)
		CheckCartExists(ctx context.Context, userId uuid.UUID) (uuid.UUID, error)
	}

	IntegrationRest interface {
		GetAllProducts() (entity.Products, error)
	}

	IntegrationContract interface {
		GetAllProducts() (entity.Products, error)
	}
)
