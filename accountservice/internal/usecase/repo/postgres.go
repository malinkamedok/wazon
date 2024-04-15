package repo

import (
	"accountservice/internal/entity"
	"accountservice/internal/usecase"
	"accountservice/pkg/postgres"
	"context"
	"github.com/Masterminds/squirrel"
	"log"
)

type PostgresRepo struct {
	*postgres.Postgres
}

func (postgres PostgresRepo) GetUserById(ctx context.Context, userId int) (entity.User, error) {
	query, args, err := postgres.Builder.
		Select("UserID", "Name", "Surname", "login").
		From("accountservice.user_account").
		Where(squirrel.Eq{"UserID": userId}).
		ToSql()
	if err != nil {
		log.Println("could not build query")
		return entity.User{}, err
	}
	log.Println(query)
	rows, err := postgres.Pool.Query(ctx, query, args...)
	var user entity.User
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&user.Id, &user.Name, &user.SurName, &user.Login)
	if err != nil {
		log.Println("could not scan row")
		log.Println(err)
		return entity.User{}, err
	}
	return user, nil
}

func (postgres PostgresRepo) InsertOrUpdateProduct(ctx context.Context, product entity.Product) error {
	query := "Insert into accountservice.product values ($1, $2, $3, $4) " +
		"on conflict (id) do update set name = EXCLUDED.name, description = EXCLUDED.description, price = EXCLUDED.price"
	log.Println(query)
	_, err := postgres.Pool.Exec(ctx, query, product.Id, product.Name, product.Description, product.Price)
	if err != nil {
		log.Println("could not execute query")
		log.Println(err)
		return err
	}
	return nil
}

func (p PostgresRepo) GetAllProducts(ctx context.Context, userId int64) ([]entity.Product, error) {
	query := "SELECT p.id, p.name, p.description, p.price FROM accountservice.product p " +
		"JOIN accountservice.product_cart pc ON p.id = pc.productID " +
		"JOIN accountservice.cart c ON pc.cartID = c.cartID " +
		"JOIN accountservice.user_account ua ON c.userID = ua.userID " +
		"WHERE ua.userID = $1;"
	rows, err := p.Pool.Query(ctx, query, userId)
	if err != nil {
		log.Println("could not execute query")
		return nil, err
	}

	var products []entity.Product

	defer rows.Close()
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price)
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

var _ usecase.AccountServiceRepository = (*PostgresRepo)(nil)

func NewPostgresRepo(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{pg}
}
