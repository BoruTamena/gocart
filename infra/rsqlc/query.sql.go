// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package rsqlc

import (
	"context"
	"database/sql"
)

const addCartItem = `-- name: AddCartItem :one
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
func (q *Queries) AddCartItem(ctx context.Context, arg AddCartItemParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, addCartItem, arg.SessionID, arg.ProductID, arg.Quantity)
	var quantity int32
	err := row.Scan(&quantity)
	return quantity, err
}

const createShoppingSession = `-- name: CreateShoppingSession :one
INSERT INTO shopping_session (user_id, total, created_at, modified_at)
VALUES ($1, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id
`

// CreateShoppingSession
//
//	INSERT INTO shopping_session (user_id, total, created_at, modified_at)
//	VALUES ($1, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
//	RETURNING id
func (q *Queries) CreateShoppingSession(ctx context.Context, userID sql.NullInt32) (int32, error) {
	row := q.db.QueryRowContext(ctx, createShoppingSession, userID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const decreaseQuantity = `-- name: DecreaseQuantity :exec
UPDATE cart_item 
SET quantity=quantity-1
WHERE session_id = $1 AND product_id = $2 AND quantity > 1
`

type DecreaseQuantityParams struct {
	SessionID sql.NullInt32 `db:"session_id" json:"session_id"`
	ProductID sql.NullInt32 `db:"product_id" json:"product_id"`
}

// DecreaseQuantity
//
//	UPDATE cart_item
//	SET quantity=quantity-1
//	WHERE session_id = $1 AND product_id = $2 AND quantity > 1
func (q *Queries) DecreaseQuantity(ctx context.Context, arg DecreaseQuantityParams) error {
	_, err := q.db.ExecContext(ctx, decreaseQuantity, arg.SessionID, arg.ProductID)
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

const increaseQuantity = `-- name: IncreaseQuantity :exec
UPDATE cart_item 
SET quantity=quantity+1
WHERE session_id = $1 AND product_id = $2
`

type IncreaseQuantityParams struct {
	SessionID sql.NullInt32 `db:"session_id" json:"session_id"`
	ProductID sql.NullInt32 `db:"product_id" json:"product_id"`
}

// IncreaseQuantity
//
//	UPDATE cart_item
//	SET quantity=quantity+1
//	WHERE session_id = $1 AND product_id = $2
func (q *Queries) IncreaseQuantity(ctx context.Context, arg IncreaseQuantityParams) error {
	_, err := q.db.ExecContext(ctx, increaseQuantity, arg.SessionID, arg.ProductID)
	return err
}

const removeCartItem = `-- name: RemoveCartItem :execrows
WITH delete_item AS (
    DELETE FROM cart_item
    WHERE session_id = $1 AND product_id = $2
    RETURNING session_id
),
updated_session AS (
    UPDATE shopping_session 
    SET total = (
        SELECT COALESCE(SUM(quantity), 0)
        FROM cart_item
        WHERE session_id = $1
    )
    WHERE id = $1
    RETURNING id, total
)
UPDATE shopping_session 
SET total = 0 
WHERE id = $1 AND NOT EXISTS (
    SELECT 1 FROM cart_item WHERE session_id = $1
)
`

type RemoveCartItemParams struct {
	Column1 sql.NullInt32 `db:"column_1" json:"column_1"`
	Column2 sql.NullInt32 `db:"column_2" json:"column_2"`
}

// RemoveCartItem
//
//	WITH delete_item AS (
//	    DELETE FROM cart_item
//	    WHERE session_id = $1 AND product_id = $2
//	    RETURNING session_id
//	),
//	updated_session AS (
//	    UPDATE shopping_session
//	    SET total = (
//	        SELECT COALESCE(SUM(quantity), 0)
//	        FROM cart_item
//	        WHERE session_id = $1
//	    )
//	    WHERE id = $1
//	    RETURNING id, total
//	)
//	UPDATE shopping_session
//	SET total = 0
//	WHERE id = $1 AND NOT EXISTS (
//	    SELECT 1 FROM cart_item WHERE session_id = $1
//	)
func (q *Queries) RemoveCartItem(ctx context.Context, arg RemoveCartItemParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, removeCartItem, arg.Column1, arg.Column2)
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

const viewCurrentCartITem = `-- name: ViewCurrentCartITem :many
SELECT id, name, description, sku, category, price, discount_id, created_at, modified_at FROM  product
WHERE product.id IN (
    SELECT product_id 
    FROM cart_item 
    WHERE session_id=$1
)
`

// ViewCurrentCartITem
//
//	SELECT id, name, description, sku, category, price, discount_id, created_at, modified_at FROM  product
//	WHERE product.id IN (
//	    SELECT product_id
//	    FROM cart_item
//	    WHERE session_id=$1
//	)
func (q *Queries) ViewCurrentCartITem(ctx context.Context, sessionID sql.NullInt32) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, viewCurrentCartITem, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Sku,
			&i.Category,
			&i.Price,
			&i.DiscountID,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
