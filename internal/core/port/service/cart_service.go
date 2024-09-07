package service

import "github.com/BoruTamena/internal/core/models"

type CartService interface {
	CreateShoppingSession() (int, error)
	AddItem(models.Item) (int, error)
	IncreaseItemQuantity(models.Item) (int, error)
	DecreaseItemQuantity(models.Item) (int, error)
	RemoveItem(models.DeletedItem) (int, error)
	UpdateQuantity(int) (models.Item, error)
	ViewCartItem(int) ([]models.Item, error)
	Checkout()
}
