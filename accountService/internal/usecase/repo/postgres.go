package repo

import (
	"accountservice/internal/entity"
	"accountservice/internal/usecase"
	"accountservice/pkg/postgres"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"log"
)

type PostgresRepo struct {
	*postgres.Postgres
}

var _ usecase.AccountServiceRepository = (*PostgresRepo)(nil)

func NewPostgresRepo(pg *postgres.Postgres) *PostgresRepo {
	return &PostgresRepo{pg}
}

func (postgres PostgresRepo) GetUserById(ctx context.Context, userId uuid.UUID) (entity.User, error) {
	query, args, err := postgres.Builder.
		Select("id", "Name", "Surname", "login").
		From("accountService.user_account").
		Where(squirrel.Eq{"id": userId}).
		ToSql()
	if err != nil {
		log.Println("could not build query")
		return entity.User{}, err
	}
	log.Println(query)
	rows, err := postgres.Pool.Query(ctx, query, args...)
	if err != nil {
		log.Println("could not execute query")
		return entity.User{}, err
	}
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
	query := "Insert into accountService.product values ($1, $2, $3, $4) " +
		"on conflict (id) do update set name = EXCLUDED.name, description = EXCLUDED.description, price = EXCLUDED.price"
	_, err := postgres.Pool.Exec(ctx, query, product.Id, product.Name, product.Description, product.Price)
	if err != nil {
		log.Println("could not execute query")
		log.Println(err)
		return err
	}
	return nil
}

func (postgres PostgresRepo) GetAllProductsFromCart(ctx context.Context, userId uuid.UUID) ([]entity.Product, error) {
	query := "SELECT p.id, p.name, p.description, p.price FROM accountService.product p " +
		"JOIN accountService.product_cart pc ON p.id = pc.productID " +
		"JOIN accountService.cart c ON pc.cartID = c.id " +
		"JOIN accountService.user_account ua ON c.userID = ua.id " +
		"WHERE ua.id = $1;"
	rows, err := postgres.Pool.Query(ctx, query, userId)
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

func (postgres PostgresRepo) CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error) {
	query, _, err := postgres.Builder.Insert("accountservice.user_account").Columns("Name", "Surname", "login").Values(user.Name, user.SurName, user.Login).Suffix("RETURNING accountservice.user_account.id").ToSql()
	if err != nil {
		log.Println("could not build query")
		return uuid.Nil, err
	}

	rows, err := postgres.Pool.Query(ctx, query, user.Name, user.SurName, user.Login)
	if err != nil {
		log.Println("could not execute query")
		return uuid.Nil, err
	}
	defer rows.Close()

	var userId uuid.UUID
	for rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			log.Println("could not scan row")
			return uuid.Nil, err
		}
	}

	return userId, err
}
