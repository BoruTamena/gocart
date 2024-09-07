// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package rsqlc

import (
	"context"
	"database/sql"
)

const addCartItem = `-- name: AddCartItem :exec
INSERT INTO cart_item (session_id, product_id, quantity, created_at, modified_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (session_id, product_id) DO UPDATE 
SET quantity = CASE WHEN cart_item.quantity <  10  THEN cart_item.quantity + EXCLUDED.quantity
    ELSE 10 
END,
modified_at = CURRENT_TIMESTAMP 
RETURNING cart_item.quantity
`

type AddCartItemParams struct {
	SessionID sql.NullInt32 `db:"session_id" json:"session_id"`
	ProductID sql.NullInt32 `db:"product_id" json:"product_id"`
	Quantity  int32         `db:"quantity" json:"quantity"`
}

// AddCartItem
//
//	INSERT INTO cart_item (session_id, product_id, quantity, created_at, modified_at)
//	VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
//	ON CONFLICT (session_id, product_id) DO UPDATE
//	SET quantity = CASE WHEN cart_item.quantity <  10  THEN cart_item.quantity + EXCLUDED.quantity
//	    ELSE 10
//	END,
//	modified_at = CURRENT_TIMESTAMP
//	RETURNING cart_item.quantity
func (q *Queries) AddCartItem(ctx context.Context, arg AddCartItemParams) error {
	_, err := q.db.ExecContext(ctx, addCartItem, arg.SessionID, arg.ProductID, arg.Quantity)
	return err
}

const createShoppingSession = `-- name: CreateShoppingSession :exec
INSERT INTO shopping_session (user_id, total, created_at, modified_at)
VALUES ($1, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id
`

// CreateShoppingSession
//
//	INSERT INTO shopping_session (user_id, total, created_at, modified_at)
//	VALUES ($1, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
//	RETURNING id
func (q *Queries) CreateShoppingSession(ctx context.Context, userID sql.NullInt32) error {
	_, err := q.db.ExecContext(ctx, createShoppingSession, userID)
	return err
}

const getActiveSession = `-- name: GetActiveSession :one
SELECT id, user_id, total, created_at, modified_at 
FROM shopping_session
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT 1
`

// GetActiveSession
//
//	SELECT id, user_id, total, created_at, modified_at
//	FROM shopping_session
//	WHERE user_id = $1
//	ORDER BY created_at DESC
//	LIMIT 1
func (q *Queries) GetActiveSession(ctx context.Context, userID sql.NullInt32) (ShoppingSession, error) {
	row := q.db.QueryRowContext(ctx, getActiveSession, userID)
	var i ShoppingSession
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Total,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const removeCartItem = `-- name: RemoveCartItem :execrows
DELETE FROM cart_item
WHERE session_id = $1 AND product_id = $2
`

type RemoveCartItemParams struct {
	SessionID sql.NullInt32 `db:"session_id" json:"session_id"`
	ProductID sql.NullInt32 `db:"product_id" json:"product_id"`
}

// RemoveCartItem
//
//	DELETE FROM cart_item
//	WHERE session_id = $1 AND product_id = $2
func (q *Queries) RemoveCartItem(ctx context.Context, arg RemoveCartItemParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, removeCartItem, arg.SessionID, arg.ProductID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const updateCartItemQuantity = `-- name: UpdateCartItemQuantity :exec
UPDATE cart_item
SET quantity = $3, modified_at = CURRENT_TIMESTAMP
WHERE session_id = $1 AND product_id = $2
RETURNING id
`

type UpdateCartItemQuantityParams struct {
	SessionID sql.NullInt32 `db:"session_id" json:"session_id"`
	ProductID sql.NullInt32 `db:"product_id" json:"product_id"`
	Quantity  int32         `db:"quantity" json:"quantity"`
}

// UpdateCartItemQuantity
//
//	UPDATE cart_item
//	SET quantity = $3, modified_at = CURRENT_TIMESTAMP
//	WHERE session_id = $1 AND product_id = $2
//	RETURNING id
func (q *Queries) UpdateCartItemQuantity(ctx context.Context, arg UpdateCartItemQuantityParams) error {
	_, err := q.db.ExecContext(ctx, updateCartItemQuantity, arg.SessionID, arg.ProductID, arg.Quantity)
	return err
}
