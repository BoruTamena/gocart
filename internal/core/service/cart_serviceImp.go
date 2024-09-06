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

func (cs cartService) AddItem(product models.Product) (int, error) {

}

func (cs cartService) RemoveItem(item_id int) error                       {}
func (cs cartService) UpdateQuantity(item_id int) (models.Product, error) {}
func (cs cartService) ViewCartItem() ([]models.Product, eror)             {}
func (cs cartService) Checkout()                                          {}
