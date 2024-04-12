package repo

import (
	"context"
	"delivery/internal/entity"
	"delivery/internal/usecase"
	"delivery/pkg/postgres"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type PostgresRepo struct {
	*postgres.Postgres
}

var _ usecase.DeliveryRepository = (*PostgresRepo)(nil)

func NewPostgresRepo(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{pg}
}

func (p PostgresRepo) ReadAllOrders(ctx context.Context) ([]entity.OrderList, error) {
	query, _, err := p.Builder.Select("id", "order_status").From("orders").ToSql()
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

func ReadOneObjectFromRows(rows pgx.Rows) (entity.Order, error) {
	var order entity.Order
	var err error

	defer rows.Close()
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

func (p PostgresRepo) ReadOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error) {
	query, _, err := p.Builder.Select("*").From("orders").Where("id = $1").ToSql()

	if err != nil {
		log.Println("could not build query")
		return entity.Order{}, err
	}
	rows, err := p.Pool.Query(ctx, query, orderUUID.String())
	if err != nil {
		log.Println("could not execute query")
		return entity.Order{}, err
	}

	order, err := ReadOneObjectFromRows(rows)

	return order, err
}

func (p *PostgresRepo) InsertOrder(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error) {
	query, _, err := p.Builder.Insert("orders").Columns("id").Values(orderUUID.String()).Suffix("RETURNING *").ToSql()

	if err != nil {
		log.Println("could not build query")
		return entity.Order{}, err
	}
	rows, err := p.Pool.Query(ctx, query, orderUUID.String())
	if err != nil {
		log.Println("could not insert into table")
		return entity.Order{}, err
	}

	order, err := ReadOneObjectFromRows(rows)

	return order, err
}

func (p *PostgresRepo) UpdateOrderByUUID(ctx context.Context, orderUUID uuid.UUID, Status entity.OrderStatus) (entity.Order, error) {
	query, _, err := p.Builder.Update("orders").Set("order_status", Status).Set("updated_at", time.Now()).Where("id = $3", orderUUID).Suffix("RETURNING *").ToSql()

	if err != nil {
		log.Println("could not build query")
		return entity.Order{}, err
	}

	rows, err := p.Pool.Query(ctx, query, Status, time.Now(), orderUUID.String())
	if err != nil {
		log.Println("could not update value")
		return entity.Order{}, err
	}

	order, err := ReadOneObjectFromRows(rows)

	return order, err
}
