package usecase

import (
	"context"
	"github.com/google/uuid"
	"log"
)

type MqUseCase struct {
	mq DeliveryMQ
}

func (m MqUseCase) SendOrderUpdateMessage(ctx context.Context, orderID string, status string) error {
	orderUUID, err := uuid.Parse(orderID)
	if err != nil {
		log.Println("error parsing order uuid")
		return err
	}

	return m.mq.SendOrderUpdateMessage(ctx, orderUUID, status)
}

func NewDeliveryMQUseCase(mq DeliveryMQ) *MqUseCase {
	return &MqUseCase{mq: mq}
}

var _ DeliveryMQContract = (*MqUseCase)(nil)
