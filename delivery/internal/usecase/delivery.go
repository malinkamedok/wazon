package usecase

import (
	"context"
	"delivery/internal/entity"

	"github.com/google/uuid"
)

type DeliveryUseCase struct {
	repo DeliveryRepository
}

var _ DeliveryContract = (*DeliveryUseCase)(nil)

func NewDeliveryUseCase(repo DeliveryRepository) *DeliveryUseCase {
	return &DeliveryUseCase{repo: repo}
}

func (s DeliveryUseCase) GetAllOrders(ctx context.Context) ([]entity.OrderList, error) {
	products, err := s.repo.ReadAllOrders(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s DeliveryUseCase) GetOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error) {
	order, err := s.repo.ReadOrderByUUID(ctx, orderUUID)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

// CreateOrder implements DeliveryContract.
func (s *DeliveryUseCase) CreateOrder(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error) {
	panic("unimplemented")
}

// UpdateOrderByUUID implements DeliveryContract.
func (s *DeliveryUseCase) UpdateOrderByUUID(ctx context.Context, orderUUID uuid.UUID, Status entity.OrderStatus) (entity.Order, error) {
	panic("unimplemented")
}
