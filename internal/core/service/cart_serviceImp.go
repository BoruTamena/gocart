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

}

func (cs cartService) RemoveItem(item_id int) error                    {}
func (cs cartService) UpdateQuantity(item_id int) (models.Item, error) {}
func (cs cartService) ViewCartItem() ([]models.Item, error)            {}
func (cs cartService) Checkout()                                       {}
