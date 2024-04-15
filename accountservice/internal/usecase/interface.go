package usecase

import (
	"accountservice/internal/entity"
	"context"
)

type (
	AccountServiceContract interface {
		GetUserById(ctx context.Context, userId int) (entity.User, error)
		InsertOrUpdateProduct(ctx context.Context, product entity.Product) error
	}

	AccountServiceRepository interface {
		GetUserById(ctx context.Context, userId int) (entity.User, error)
		InsertOrUpdateProduct(ctx context.Context, product entity.Product) error
	}
)
