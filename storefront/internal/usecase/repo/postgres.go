package repo

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"storefront/internal/entity"
	"storefront/internal/usecase"
	"storefront/pkg/logger"
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
		logger.Error("could not build query", zap.Error(err))
		return nil, err
	}
	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		logger.Error("could not execute query", zap.Error(err))
		return nil, err
	}

	var products []entity.ProductList

	defer rows.Close()
	for rows.Next() {
		var product entity.ProductList
		err = rows.Scan(&product.UUID, &product.Name)
		if err != nil {
			logger.Error("could not scan row", zap.Error(err))
			return nil, err
		}
		products = append(products, product)
	}
	if rows.Err() != nil {
		logger.Error("could not read rows", zap.Error(rows.Err()))
		return nil, err
	}
	return products, nil
}

func (p PostgresRepo) ReadProductByUUID(ctx context.Context, productUUID uuid.UUID) (entity.Product, error) {
	query, _, err := p.Builder.Select("*").From("products").Where("id=?", productUUID).ToSql()
	if err != nil {
		logger.Error("could not build query", zap.Error(err))
		return entity.Product{}, err
	}
	rows, err := p.Pool.Query(ctx, query, productUUID)
	if err != nil {
		logger.Error("could not execute query", zap.Error(err))
		return entity.Product{}, err
	}

	var product entity.Product

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.UUID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			logger.Error("could not scan row", zap.Error(err))
			return entity.Product{}, err
		}
	}
	if rows.Err() != nil {
		logger.Error("could not read rows", zap.Error(rows.Err()))
		return entity.Product{}, err
	}
	return product, nil
}
