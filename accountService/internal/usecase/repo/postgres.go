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

func (postgres PostgresRepo) GetAllProductsFromCart(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, error) {
	query := "select productID from accountservice.product_cart " +
		"join accountservice.cart on product_cart.cartID = cart.id " +
		"where cart.userID = $1;"

	rows, err := postgres.Pool.Query(ctx, query, userId)
	if err != nil {
		log.Println("could not execute query")
		return nil, err
	}

	var products []uuid.UUID

	defer rows.Close()
	for rows.Next() {
		var product uuid.UUID
		err = rows.Scan(&product)
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

func (postgres PostgresRepo) AddProductToCart(ctx context.Context, cartID uuid.UUID, productId uuid.UUID) error {
	query, _, err := postgres.Builder.Insert("accountservice.product_cart").Columns("cartID", "productID").Values(cartID, productId).ToSql()
	if err != nil {
		log.Println("could not build query")
		return err
	}

	_, err = postgres.Pool.Exec(ctx, query, cartID, productId)
	if err != nil {
		log.Println("could not execute query")
		return err
	}

	return nil
}

func (postgres PostgresRepo) CreateCart(ctx context.Context, userId uuid.UUID) (uuid.UUID, error) {
	query, _, err := postgres.Builder.Insert("accountservice.cart").Columns("userID").Values(userId).Suffix("RETURNING accountservice.cart.id").ToSql()
	if err != nil {
		log.Println("could not build query")
		return uuid.Nil, err
	}

	rows, err := postgres.Pool.Query(ctx, query, userId)
	if err != nil {
		log.Println("could not execute query")
		return uuid.Nil, err
	}
	defer rows.Close()

	var cartId uuid.UUID
	for rows.Next() {
		err = rows.Scan(&cartId)
		if err != nil {
			log.Println("could not scan row")
			return uuid.Nil, err
		}
	}
	return cartId, nil
}

func (postgres PostgresRepo) CheckCartExists(ctx context.Context, userId uuid.UUID) (uuid.UUID, error) {
	query, _, err := postgres.Builder.Select("id").From("accountservice.cart").Where(squirrel.Eq{"userID": userId}).ToSql()

	rows, err := postgres.Pool.Query(ctx, query, userId)
	if err != nil {
		log.Println("could not execute query")
		return uuid.Nil, err
	}
	defer rows.Close()

	var cartID uuid.UUID
	for rows.Next() {
		err = rows.Scan(&cartID)
		if err != nil {
			log.Println("could not scan row")
			return uuid.Nil, err
		}
	}

	return cartID, nil
}
