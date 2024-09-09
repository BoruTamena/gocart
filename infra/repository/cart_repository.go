package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/BoruTamena/infra/rsqlc"
	"github.com/BoruTamena/internal/core/port/repository"
)

type cartRepository struct {
	db repository.DataBase
}

func NewCartRepository(database repository.DataBase) repository.CartRepository {
	return &cartRepository{
		db: database,
	}
}

// full-stack
func (cr cartRepository) InsertShoppingSession(c context.Context, user_id sql.NullInt32) (int, error) {

	query := rsqlc.New(cr.db.GetDB())

	// defer cr.db.Close()

	session_id, err := query.CreateShoppingSession(c, user_id)

	if err != nil {
		return 0, err
	}

	return int(session_id), nil

}

// full-stack
func (cr cartRepository) InserCartItem(c context.Context, item_param rsqlc.AddCartItemParams) (int, error) {

	query := rsqlc.New(cr.db.GetDB())

	// defer cr.db.Close()

	quantity, err := query.AddCartItem(c, item_param)

	if err != nil {
		return 0, err
	}

	return int(quantity), nil

}

func (cr cartRepository) IncreaseQuantity(c context.Context, quantity_param rsqlc.IncreaseQuantityParams) error {

	query := rsqlc.New(cr.db.GetDB())

	// defer cr.db.Close()

	err := query.IncreaseQuantity(c, quantity_param)
	if err != nil {
		return err
	}

	return nil

}

func (cr cartRepository) DecreaseQuantity(c context.Context, quantity_param rsqlc.DecreaseQuantityParams) error {
	query := rsqlc.New(cr.db.GetDB())

	// defer cr.db.Close()

	err := query.DecreaseQuantity(c, quantity_param)
	if err != nil {
		return err
	}

	return nil
}

// full-stack
func (cr cartRepository) DeleteCartItem(c context.Context, item_param rsqlc.RemoveCartItemParams) (int, error) {

	query := rsqlc.New(cr.db.GetDB())

	// defer cr.db.Close()

	affected_row, err := query.RemoveCartItem(c, item_param)

	if err != nil {
		return 0, err
	}

	return int(affected_row), nil
}

func (cr cartRepository) SelectCartItem(c context.Context, session_id sql.NullInt32) ([]rsqlc.Product, error) {

	query := rsqlc.New(cr.db.GetDB())

	defer cr.db.Close()

	// product managers
	product, err := query.ViewCurrentCartItem(c, session_id)

	if err != nil {

		log.Println("error is come from here...", err)
		return nil, err
	}

	return product, err

}

func (cr cartRepository) CartCheckOut(c context.Context) {}
