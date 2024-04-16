package repo

import (
	"context"
	"github.com/google/uuid"
	"log"
	"storefront/internal/entity"
	"storefront/internal/usecase"
	"storefront/pkg/postgres"
)

type PostgresRepo struct {
	*postgres.Postgres
}

var _ usecase.StorefrontRepository = (*PostgresRepo)(nil)

func NewPostgresRepo(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{pg}
}

func (p PostgresRepo) ReadAllProducts(ctx context.Context) ([]entity.ProductList, error) {
	query, _, err := p.Builder.Select("id", "name").From("productcard.products").ToSql()
	if err != nil {
		log.Println("could not build query")
		return nil, err
	}
	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		log.Println("could not execute query")
		return nil, err
	}

	var products []entity.ProductList

	defer rows.Close()
	for rows.Next() {
		var product entity.ProductList
		err = rows.Scan(&product.UUID, &product.Name)
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

func (p PostgresRepo) ReadProductByUUID(ctx context.Context, productUUID uuid.UUID) (entity.Product, error) {
	//TODO implement me
	panic("implement me")
}
