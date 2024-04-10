package repo

import (
	"context"
	"delivery/internal/entity"
	"delivery/internal/usecase"
	"delivery/pkg/postgres"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type PostgresRepo struct {
	*postgres.Postgres
}

var _ usecase.DeliveryRepository = (*PostgresRepo)(nil)

func NewPostgresRepo(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{pg}
}

func (p PostgresRepo) ReadAllOrders(ctx context.Context) ([]entity.OrderList, error) {
	query, _, err := p.Builder.Select("id", "name").From("products").ToSql()
	if err != nil {
		log.Println("could not build query")
		return nil, err
	}
	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		log.Println("could not execute query")
		return nil, err
	}

	var products []entity.OrderList

	defer rows.Close()
	for rows.Next() {
		var product entity.OrderList
		err = rows.Scan(&product.UUID, &product.Status)
		if err != nil {
			log.Println("could not scan row")
			return nil, err
		}
		products = append(products, product)
	}
	if rows.Err() != nil {
		log.Println("could not read rows")
		return nil, err
	}
	return products, nil
}

func (p PostgresRepo) ReadOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error) {
	query, _, err := p.Builder.Select("id", "status").From("orders").Where("id", "=", orderUUID).ToSql()

	if err != nil {
		log.Println("could not build query")
		return entity.Order{}, err
	}
	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		log.Println("could not execute query")
		return entity.Order{}, err
	}
	defer rows.Close()
	var order entity.Order
	for rows.Next() {
		err = rows.Scan(&order.UUID, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			log.Println("could not scan row")
			return entity.Order{}, err
		}
	}
	if rows.Err() != nil {
		log.Println("could not read rows")
		return entity.Order{}, err
	}
	if order == (entity.Order{}) {
		return entity.Order{}, fmt.Errorf("no orders were obtained")
	}
	return order, nil
}

func (p *PostgresRepo) InsertOrder(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error) {
	panic("unimplemented")
}

func (p *PostgresRepo) UpdateOrderByUUID(ctx context.Context, orderUUID uuid.UUID, Status entity.OrderStatus) (entity.Order, error) {
	panic("unimplemented")
}
