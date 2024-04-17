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

func (p *PostgresRepo) ReadAllOrders(ctx context.Context) ([]entity.OrderList, error) {
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

	for rows.Next() {
		err = rows.Scan(&order.UUID, &order.Status, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			log.Println("could not scan row")
			return entity.Order{}, err
		}
	}
	return order, nil
}

func (p *PostgresRepo) ReadOrderByUUID(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error) {
	query, _, err := p.Builder.Select("*").From("order_service.orders").Where("id = $1").ToSql()

	if err != nil {
		log.Println("could not build query")
		return entity.Order{}, err
	}
	rows, err := p.Pool.Query(ctx, query, orderUUID.String())
	if err != nil {
		log.Println("could not execute query")
		return entity.Order{}, err
	}
	defer rows.Close()

	order, err := ReadOneObjectFromRows(rows)

	return order, err
}

func (p *PostgresRepo) InsertOrder(ctx context.Context, orderUUID uuid.UUID) (entity.Order, error) {
	exists, err := p.CheckOrderExistance(ctx, orderUUID)

	if err != nil {
		return entity.Order{}, err
	}

	if exists {
		return entity.Order{}, fmt.Errorf("order with UUID %s already exists", orderUUID)
	}

	query, _, err := p.Builder.Insert("order_service.orders").Columns("id").Values(orderUUID.String()).Suffix("RETURNING *").ToSql()

	if err != nil {
		log.Println("could not build query")
		return entity.Order{}, err
	}
	rows, err := p.Pool.Query(ctx, query, orderUUID.String())
	if err != nil {
		log.Println("could not insert into table")
		return entity.Order{}, err
	}
	defer rows.Close()

	order, err := ReadOneObjectFromRows(rows)

	return order, err
}

func (p *PostgresRepo) UpdateOrderByUUID(ctx context.Context, orderUUID uuid.UUID, Status entity.OrderStatus) (entity.Order, error) {
	exists, err := p.CheckOrderExistance(ctx, orderUUID)

	if err != nil {
		return entity.Order{}, err
	}

	if !exists {
		return entity.Order{}, fmt.Errorf("order with UUID %s does not exist yet", orderUUID)
	}

	query, _, err := p.Builder.Update("order_service.orders").Set("order_status", Status).Set("updated_at", time.Now()).Where("id = $3", orderUUID).Suffix("RETURNING *").ToSql()

	if err != nil {
		log.Println("could not build query")
		return entity.Order{}, err
	}

	rows, err := p.Pool.Query(ctx, query, Status, time.Now(), orderUUID.String())
	if err != nil {
		log.Println("could not update value")
		return entity.Order{}, err
	}
	defer rows.Close()

	order, err := ReadOneObjectFromRows(rows)

	return order, err
}

func (p *PostgresRepo) CheckOrderExistance(ctx context.Context, orderUUID uuid.UUID) (bool, error) {
	var exists bool
	query, _, err := p.Builder.Select("1").Prefix("SELECT EXISTS (").From("order_service.orders").Where("id = $1", orderUUID).Suffix(")").ToSql()

	if err != nil {
		log.Println("could not build query")
		return false, err
	}

	err = p.Pool.QueryRow(ctx, query, orderUUID).Scan(&exists)
	if err != nil {
		log.Println("could not check for existanse")
		return false, err
	}
	return exists, nil
}
