package service

import "github.com/BoruTamena/internal/core/models"

type CartService interface {
	AddItem(models.Product) (int, error)
	RemoveItem(int) error
	UpdateQuantity(int) (models.Product, error)
	ViewCartItem() ([]models.Product, error)
	Checkout()
}
