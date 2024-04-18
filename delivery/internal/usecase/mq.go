package usecase

import "context"

type MqUseCase struct {
	mq DeliveryMQ
}

func (m MqUseCase) SendOrderUpdateMessage(ctx context.Context, orderUUID string, status string) error {
	//TODO implement me
	panic("implement me")
}

func NewDeliveryMQUseCase(mq DeliveryMQ) *MqUseCase {
	return &MqUseCase{mq: mq}
}

var _ DeliveryMQContract = (*MqUseCase)(nil)
