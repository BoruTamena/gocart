package service

import (
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

func (cs cartService) AddItem()        {}
func (cs cartService) RemoveItem()     {}
func (cs cartService) UpdateQuantity() {}
func (cs cartService) ViewCartItem()   {}
func (cs cartService) Checkout()       {}
