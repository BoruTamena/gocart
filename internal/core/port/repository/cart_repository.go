package repository

import (
	"context"
	"database/sql"

	"github.com/BoruTamena/infra/rsqlc"
)

type CartRepository interface {
	InsertShoppingSession(context.Context, sql.NullInt32) (int, error) // return session_id and error if any
	InserCartItem(context.Context, rsqlc.AddCartItemParams) (int, error)
	IncreaseQuantity(context.Context, rsqlc.IncreaseQuantityParams) error
	DecreaseQuantity(context.Context, rsqlc.DecreaseQuantityParams) error
	DeleteCartItem(context.Context, rsqlc.RemoveCartItemParams) (int, error)
	SelectCartItem(context.Context, sql.NullInt32) ([]rsqlc.Product, error)
	CartCheckOut(context.Context, sql.NullInt32) error
}
