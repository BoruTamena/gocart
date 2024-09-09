package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/BoruTamena/infra/rsqlc"
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

	session_id, err := cs.Rep.InsertShoppingSession(c, UserIDNull)

	return session_id, err
}

func (cs cartService) AddItem(c context.Context, product models.Item) (int, error) {
	// adding new item to cart
	pr := rsqlc.AddCartItemParams{
		SessionID: sql.NullInt32{Int32: int32(product.SessionId), Valid: true},
		ProductID: sql.NullInt32{Int32: int32(product.ProductId), Valid: true},
		Quantity:  1,
	}

	quantity, err := cs.Rep.InserCartItem(c, pr)

	if err != nil {
		return 0, err
	}

	if quantity == 10 {
		return quantity, errors.New("max limit exceeded")

	}

	return quantity, nil

}
func (cs cartService) IncreaseItemQuantity(c context.Context, product models.Item) (int, error) {

	// increase existing cart item quantity
	quantity_param := rsqlc.IncreaseQuantityParams{

		SessionID: sql.NullInt32{Int32: int32(product.SessionId), Valid: true},
		ProductID: sql.NullInt32{Int32: int32(product.ProductId), Valid: true},
	}

	if err := cs.Rep.IncreaseQuantity(c, quantity_param); err != nil {
		return 0, err

	}

	return 1, nil

}
func (cs cartService) DecreaseItemQuantity(c context.Context, product models.Item) (int, error) {

	// decrease existing cart item quantity

	quantity_param := rsqlc.DecreaseQuantityParams{

		SessionID: sql.NullInt32{Int32: int32(product.SessionId), Valid: true},
		ProductID: sql.NullInt32{Int32: int32(product.ProductId), Valid: true},
	}

	if err := cs.Rep.DecreaseQuantity(c, quantity_param); err != nil {
		return 0, err

	}

	return 1, nil

}

func (cs cartService) RemoveItem(c context.Context, item models.DeletedItem) (int, error) {

	// removing cart item

	item_param := rsqlc.RemoveCartItemParams{

		Column1: sql.NullInt32{Int32: int32(item.ProductId), Valid: true},
		Column2: sql.NullInt32{Int32: int32(item.SessionId), Valid: true},
	}

	num_row, err := cs.Rep.DeleteCartItem(c, item_param)

	if err != nil {

		return 0, nil

	}

	return num_row, nil

}

func (cs cartService) ViewCartItem(c context.Context, session_id int) ([]rsqlc.Product, error) {

	/*
	 get current item

	*/
	id := sql.NullInt32{Int32: int32(session_id), Valid: true}

	log.Println("session Id", id)
	product, err := cs.Rep.SelectCartItem(c, id)

	if err != nil {
		log.Println("Db:", err)
		return nil, err
	}

	return product, nil

}
func (cs cartService) Checkout(c context.Context) {}
