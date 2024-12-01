package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/freeloginname/home_work_basic/hw14_sql/internal/repository/product"
	"github.com/freeloginname/home_work_basic/hw14_sql/pkg/pgdb"
	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	// if !ok {
	// 	fmt.Println("DB_DSN is not set")
	// 	return
	// }
	if len(dsn) == 0 {
		fmt.Println("DB_DSN is not set")
		return
	}
	var numeric pgtype.Numeric

	ctx := context.Background()
	dbc, err := pgdb.New(ctx, dsn, 5)
	if err != nil {
		fmt.Printf("failed to connect to DB: %v", err)
		return
	}
	defer dbc.Close()
	requestor := product.New(dbc)
	user1ID, err := requestor.CreateUser(ctx, product.CreateUserParams{
		Name:     "John Doe",
		Email:    "xk0e5@example.com",
		Password: "password",
	})
	if err != nil {
		fmt.Printf("failed to create user: %v", err)
		return
	}
	_, err = requestor.CreateUser(ctx, product.CreateUserParams{
		Name:     "aaaa Doe",
		Email:    "47h5@example.com",
		Password: "password",
	})
	if err != nil {
		fmt.Printf("failed to create user: %v", err)
		return
	}

	if err = numeric.Scan("10"); err != nil {
		fmt.Printf("failed to scan numeric: %v", err)
		return
	}
	_, err = requestor.CreateProduct(ctx, product.CreateProductParams{
		Name:  "Product 1",
		Price: numeric,
	})
	if err != nil {
		fmt.Printf("failed to create product: %v", err)
		return
	}

	if err = numeric.Scan("20"); err != nil {
		fmt.Printf("failed to scan numeric: %v", err)
		return
	}
	_, err = requestor.CreateProduct(ctx, product.CreateProductParams{
		Name:  "Product 2",
		Price: numeric,
	})
	if err != nil {
		fmt.Printf("failed to create product: %v", err)
		return
	}

	if err = numeric.Scan("30"); err != nil {
		fmt.Printf("failed to scan numeric: %v", err)
		return
	}
	_, err = requestor.CreateOrder(ctx, product.CreateOrderParams{
		UserID: user1ID,
		OrderDate: pgtype.Timestamptz{
			Time:             time.Now(),
			InfinityModifier: 0,
			Valid:            true,
		},
		TotalAmount: numeric,
	})
	if err != nil {
		fmt.Printf("failed to create order: %v", err)
		return
	}

	if err = numeric.Scan("40"); err != nil {
		fmt.Printf("failed to scan numeric: %v", err)
		return
	}
	_, err = requestor.CreateOrder(ctx, product.CreateOrderParams{
		UserID: user1ID,
		OrderDate: pgtype.Timestamptz{
			Time:             time.Now(),
			InfinityModifier: 0,
			Valid:            true,
		},
		TotalAmount: numeric,
	})
	if err != nil {
		fmt.Printf("failed to create order: %v", err)
		return
	}

	user1TotalPrice, err := requestor.GetUserTotalPrice(ctx, user1ID)
	if err != nil {
		fmt.Printf("failed to get user total price: %v", err)
		return
	}
	fmt.Printf("User 1 total price: %v\n", user1TotalPrice)
}
