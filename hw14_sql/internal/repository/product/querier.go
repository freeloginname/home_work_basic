// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package product

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) (uuid.UUID, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (uuid.UUID, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error)
	DeleteOrder(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	GetProductById(ctx context.Context, id uuid.UUID) (*Product, error)
	GetUserAvaragePrice(ctx context.Context, userID uuid.UUID) (float64, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserTotalPrice(ctx context.Context, userID uuid.UUID) (int64, error)
	GetordersByUser(ctx context.Context, userID uuid.UUID) ([]*Order, error)
	SordUserById(ctx context.Context, limit int32) ([]*User, error)
	SortProductByPriceASC(ctx context.Context, limit int32) ([]*Product, error)
	SortProductByPriceDESC(ctx context.Context, limit int32) ([]*Product, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (uuid.UUID, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (uuid.UUID, error)
}

var _ Querier = (*Queries)(nil)
