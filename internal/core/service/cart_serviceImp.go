package service

import (
	"context"
	"database/sql"

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

func (cs cartService) CreateShoppingSession(c context.Context, user_session models.Session) (int, error) {

	var UserIDNull sql.NullInt32

	UserIDNull = sql.NullInt32{Int32: int32(user_session.UserID), Valid: true}

	err := cs.Rep.InsertShoppingSession(c, UserIDNull)

	return user_session.UserID, err
}

func (cs cartService) AddItem(c context.Context, product models.Item) (int, error) {

	// adding new item to cart
}
func (cs cartService) IncreaseItemQuantity(c context.Context, product models.Item) (int, error) {

	// increase existing cart item quantity
}
func (cs cartService) DecreaseItemQuantity(c context.Context, product models.Item) (int, error) {

	// decrease existing cart item quantity

}
func (cs cartService) RemoveItem(c context.Context, item models.DeletedItem) (int, error) {}

func (cs cartService) ViewCartItem(c context.Context, user_id int) ([]models.Item, error) {

	// get active session of the user

	// get current item

}
func (cs cartService) Checkout(c context.Context) {}
