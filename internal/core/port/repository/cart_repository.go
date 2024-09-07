package repository

import (
	"context"
	"database/sql"

	"github.com/BoruTamena/infra/rsqlc"
)

type CartRepository interface {
	InsertShoppingSession(context.Context, sql.NullInt32) error // return session_id and error if any
	InserCartItem(context.Context, rsqlc.AddCartItemParams) error
	IncreaseQuantity(context.Context, rsqlc.IncreaseQuantityParams) error
	DecreaseQuantity(context.Context, rsqlc.DecreaseQuantityParams) error
	DeleteCartItem(context.Context, rsqlc.RemoveCartItemParams) error
	SelectCartItem(context.Context, sql.NullInt32) ([]rsqlc.Product, error)
	CartCheckOut(context.Context)
}
