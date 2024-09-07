package service

import (
	"github.com/BoruTamena/internal/core/models"
	"github.com/BoruTamena/internal/core/port/repository"
	"github.com/BoruTamena/internal/core/port/service"
)

type cartService struct {
	Rep repository.CartRepository
}

func NewCartService(Rep repository.CartRepository) service.CartService {

	return &cartService{
		Rep,
	}

}

func (cs cartService) CreateShoppingSession() (int, error) {}
func (cs cartService) AddItem(product models.Item) (int, error) {

	// adding new item to cart
}
func (cs cartService) IncreaseItemQuantity(product models.Item) (int, error) {

	// increase existing cart item quantity
}
func (cs cartService) DecreaseItemQuantity(product models.Item) (int, error) {

	// decrease existing cart item quantity

}
func (cs cartService) RemoveItem(item models.DeletedItem) (int, error) {}
func (cs cartService) UpdateQuantity(item_id int) (models.Item, error) {}
func (cs cartService) ViewCartItem(user_id int) ([]models.Item, error) {

	// get active session of the user

	// get current item

}
func (cs cartService) Checkout() {}
