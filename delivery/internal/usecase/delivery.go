package usecase

import (
	"context"
	"delivery/internal/entity"
	"fmt"

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

func (s DeliveryUseCase) GetOrderByUUID(ctx context.Context, orderUUID string) (entity.Order, error) {
	parsedUUID, err := uuid.Parse(orderUUID)
	if err != nil {
		return entity.Order{}, fmt.Errorf("cannot parse uuid %s", orderUUID)
	}
	order, err := s.repo.ReadOrderByUUID(ctx, parsedUUID)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (s *DeliveryUseCase) CreateOrder(ctx context.Context, orderUUID string) (entity.Order, error) {
	parsedUUID, err := uuid.Parse(orderUUID)
	if err != nil {
		return entity.Order{}, fmt.Errorf("cannot parse uuid %s", orderUUID)
	}
	order, err := s.repo.InsertOrder(ctx, parsedUUID)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (s *DeliveryUseCase) UpdateOrderByUUID(ctx context.Context, orderUUID string, newStatus string) (entity.Order, error) {
	status, err := entity.StringToStatus(newStatus)
	if err != nil {
		return entity.Order{}, err
	}

	parsedUUID, err := uuid.Parse(orderUUID)
	if err != nil {
		return entity.Order{}, fmt.Errorf("cannot parse uuid %s", orderUUID)
	}

	order, err := s.repo.UpdateOrderByUUID(ctx, parsedUUID, status)
	if err != nil {
		return entity.Order{}, err
	}
	return order, nil
}
