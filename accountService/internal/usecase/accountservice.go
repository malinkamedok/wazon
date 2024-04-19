package usecase

import (
	"accountservice/internal/entity"
	"context"
	"github.com/google/uuid"
	"log"
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

func (useCase AccountServiceUseCase) GetAllProductsFromCart(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, error) {
	return useCase.repo.GetAllProductsFromCart(ctx, userId)
}

func (useCase AccountServiceUseCase) InsertOrUpdateProduct(ctx context.Context, product entity.Product) error {
	return useCase.repo.InsertOrUpdateProduct(ctx, product)
}

func (useCase AccountServiceUseCase) GetUserById(ctx context.Context, userId uuid.UUID) (entity.User, error) {
	return useCase.repo.GetUserById(ctx, userId)
}

func (useCase AccountServiceUseCase) CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error) {
	return useCase.repo.CreateUser(ctx, user)
}

func (useCase AccountServiceUseCase) AddProductToCart(ctx context.Context, userId string, productId string) error {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		log.Println("error in parsing user uuid", err)
		return err
	}

	productUUID, err := uuid.Parse(productId)
	if err != nil {
		log.Println("error in parsing product uuid", err)
		return err
	}

	cartID, err := useCase.repo.CheckCartExists(ctx, userUUID)
	if err != nil {
		return err
	}

	var newCartID uuid.UUID
	if cartID == uuid.Nil {
		newCartID, err = useCase.repo.CreateCart(ctx, userUUID)
		if err != nil {
			return err
		}
	} else {
		newCartID = cartID
	}

	err = useCase.repo.AddProductToCart(ctx, newCartID, productUUID)
	if err != nil {
		return err
	}

	return nil
}
