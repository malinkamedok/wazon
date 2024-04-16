package usecase

import (
	"accountservice/internal/entity"
	"context"
	"github.com/google/uuid"
)

type AccountServiceUseCase struct {
	repo AccountServiceRepository
}

var _ AccountServiceContract = (*AccountServiceUseCase)(nil)

func NewStorefrontUseCase(repo AccountServiceRepository) *AccountServiceUseCase {
	return &AccountServiceUseCase{repo: repo}
}

func (useCase AccountServiceUseCase) GetProductByUUID(ctx context.Context, productUUID uuid.UUID) (entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (useCase AccountServiceUseCase) GetAllProductsFromCart(ctx context.Context, userId uuid.UUID) ([]entity.Product, error) {
	return useCase.repo.GetAllProductsFromCart(ctx, userId)
}

func (useCase AccountServiceUseCase) InsertOrUpdateProduct(ctx context.Context, product entity.Product) error {
	return useCase.repo.InsertOrUpdateProduct(ctx, product)
}

func (useCase AccountServiceUseCase) GetUserById(ctx context.Context, userId uuid.UUID) (entity.User, error) {
	return useCase.repo.GetUserById(ctx, userId)
}
