package transaction

import (
	"context"
	"errors"
	"fmt"

	"github.com/freeloginname/home_work_basic/hw15_go_sql/internal/repository/product"
	"github.com/freeloginname/home_work_basic/hw15_go_sql/pkg/pgdb"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetUser(ctx context.Context, dsn string, name string) (product.User, error) {
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return product.User{}, err
	}
	defer dbc.Close()
	requestor := product.New(dbc)
	user, err := requestor.GetUserByName(ctx, name)
	if err != nil {
		fmt.Printf("failed to get user by name: %v", err)
		return product.User{}, err
	}
	return *user, nil
}

func CreateUser(ctx context.Context, dsn string, name string, email string, password string) (string, error) {
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return "", err
	}
	defer dbc.Close()

	tx, err := dbc.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		fmt.Printf("failed to start transaction: %v", err)
		return "", err
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	requestor := product.New(tx)
	requestor.WithTx(tx)
	users, err := requestor.GetAllUsers(ctx)
	if err != nil {
		fmt.Printf("failed to get all users: %v", err)
		return "", err
	}
	for _, user := range users {
		if user.Name == name {
			errorString := fmt.Sprintf("user with name %v already existst with id: %v", user.Name, user.ID)
			fmt.Println(errorString)
			err = errors.New(errorString)
			return "", err
		}
	}
	userID, err := requestor.CreateUser(ctx, product.CreateUserParams{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		fmt.Printf("failed to create user: %v", err)
		return "", err
	}
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Printf("failed to commit transaction: %v", err)
		return "", err
	}
	return userID.String(), nil
}

func GetOrdersByUser(ctx context.Context, dsn string, userName string) ([]*product.Order, error) {
	var userID uuid.UUID
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return nil, err
	}
	defer dbc.Close()

	tx, err := dbc.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		fmt.Printf("failed to start transaction: %v", err)
		return []*product.Order{}, err
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	requestor := product.New(tx)
	requestor.WithTx(tx)
	users, err := requestor.GetAllUsers(ctx)
	if err != nil {
		fmt.Printf("failed to get information about users: %v", err)
		return []*product.Order{}, err
	}

	for _, user := range users {
		if user.Name == userName {
			userID = user.ID
			break
		}
	}
	if userID == [16]byte{} {
		errorString := fmt.Sprintf("user with name %v not found", userName)
		fmt.Println(errorString)
		err = errors.New(errorString)
		return []*product.Order{}, err
	}

	orders, err := requestor.GetOrdersByUser(ctx, userID)
	if err != nil {
		fmt.Printf("failed to get information about orders: %v", err)
		return []*product.Order{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Printf("failed to commit transaction: %v", err)
		return []*product.Order{}, err
	}
	return orders, nil
}

func CreateOrder(ctx context.Context, dsn string, userName string, totalAmount string) error {
	var numeric pgtype.Numeric
	numeric.Scan(totalAmount)
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return err
	}
	defer dbc.Close()
	tx, err := dbc.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		fmt.Printf("failed to start transaction: %v", err)
		return err
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	requestor := product.New(tx)
	requestor.WithTx(tx)

	users, err := requestor.GetAllUsers(ctx)
	if err != nil {
		fmt.Printf("failed to get all users: %v", err)
		return err
	}
	for _, user := range users {
		if user.Name == userName {
			_, err = requestor.CreateOrderWithCurrentDate(ctx, product.CreateOrderWithCurrentDateParams{
				UserID:      user.ID,
				TotalAmount: numeric,
			})
			if err != nil {
				fmt.Printf("failed to create order: %v", err)
				return err
			}
			err = tx.Commit(ctx)
			if err != nil {
				fmt.Printf("failed to commit transaction: %v", err)
				return err
			}
			return nil
		}
	}
	errorString := fmt.Sprintf("user with name %v not found", userName)
	fmt.Println(errorString)
	err = errors.New(errorString)
	return err
}

func GetProductByName(ctx context.Context, dsn string, name string) (product.Product, error) {
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return product.Product{}, err
	}
	defer dbc.Close()
	requestor := product.New(dbc)
	product, err := requestor.GetProductByName(ctx, name)
	if err != nil {
		fmt.Printf("failed to get product by name: %v", err)
		return *product, err
	}
	return *product, nil
}

func CreateProduct(ctx context.Context, dsn string, name string, price string) (string, error) {
	var numeric pgtype.Numeric
	numeric.Scan(price)
	dbc, err := pgdb.New(ctx, dsn, 1)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return "", err
	}
	defer dbc.Close()

	tx, err := dbc.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		fmt.Printf("failed to start transaction: %v", err)
		return "", err
	}
	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	requestor := product.New(tx)
	requestor.WithTx(tx)
	// TODO добавить проверку наличия данного товара перед созданием
	products, err := requestor.GetAllProducts(ctx)
	if err != nil {
		fmt.Printf("failed to get all products: %v", err)
		return "", err
	}
	for _, product := range products {
		if product.Name == name {
			errorString := fmt.Sprintf("product with name %v already existst with id: %v", product.Name, product.ID)
			fmt.Println(errorString)
			err = errors.New(errorString)
			return "", err
		}
	}
	productID, err := requestor.CreateProduct(ctx, product.CreateProductParams{
		Name:  name,
		Price: numeric,
	})
	if err != nil {
		fmt.Printf("failed to create product: %v", err)
		return "", err
	}
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Printf("failed to commit transaction: %v", err)
		return "", err
	}
	return productID.String(), nil
}
