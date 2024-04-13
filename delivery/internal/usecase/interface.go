package usecase

import (
	"context"
	"delivery/internal/entity"

	"github.com/google/uuid"
)

type (
	DeliveryContract interface {
		GetAllOrders(ctx context.Context) ([]entity.OrderList, error)
		GetOrderByUUID(ctx context.Context, orderUUID string) (entity.Order, error)

		CreateOrder(ctx context.Context, orderUUID string) (entity.Order, error)
		UpdateOrderByUUID(ctx context.Context, orderUUID string, newStatus string) (entity.Order, error)
	}

	DeliveryRepository interface {
		ReadAllOrders(ctx context.Context) ([]entity.OrderList, error)
		ReadOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error)

		InsertOrder(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error)
		UpdateOrderByUUID(ctx context.Context, orderUUID uuid.UUID, Status entity.OrderStatus) (entity.Order, error)

		CheckOrderExistance(ctx context.Context, orderUUID uuid.UUID) (bool, error)
	}
)
