package service

import (
	"context"

	"github.com/BoruTamena/infra/rsqlc"
	"github.com/BoruTamena/internal/core/models"
)

type CartService interface {
	CreateShoppingSession(context.Context, models.Session) (int, error)
	AddItem(context.Context, models.Item) (int, error)
	IncreaseItemQuantity(context.Context, models.Item) (int, error)
	DecreaseItemQuantity(context.Context, models.Item) (int, error)
	RemoveItem(context.Context, models.DeletedItem) (int, error)
	ViewCartItem(context.Context, int) ([]rsqlc.Product, error)
	Checkout(context.Context)
}
