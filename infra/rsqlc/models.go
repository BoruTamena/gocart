// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package rsqlc

import (
	"database/sql"
)

type CartItem struct {
	ID         int32         `db:"id" json:"id"`
	SessionID  sql.NullInt32 `db:"session_id" json:"session_id"`
	ProductID  sql.NullInt32 `db:"product_id" json:"product_id"`
	Quantity   int32         `db:"quantity" json:"quantity"`
	CreatedAt  sql.NullTime  `db:"created_at" json:"created_at"`
	ModifiedAt sql.NullTime  `db:"modified_at" json:"modified_at"`
}

type Discount struct {
	ID          int32          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
	Percent     int32          `db:"percent" json:"percent"`
	CreatedAt   sql.NullTime   `db:"created_at" json:"created_at"`
	ModifiedAt  sql.NullTime   `db:"modified_at" json:"modified_at"`
}

type OrderDetails struct {
	ID         int32          `db:"id" json:"id"`
	UserID     sql.NullInt32  `db:"user_id" json:"user_id"`
	Total      sql.NullString `db:"total" json:"total"`
	PaymentID  sql.NullInt32  `db:"payment_id" json:"payment_id"`
	CreatedAt  sql.NullTime   `db:"created_at" json:"created_at"`
	ModifiedAt sql.NullTime   `db:"modified_at" json:"modified_at"`
}

type OrderItems struct {
	ID         int32         `db:"id" json:"id"`
	OrderID    sql.NullInt32 `db:"order_id" json:"order_id"`
	ProductID  sql.NullInt32 `db:"product_id" json:"product_id"`
	CreatedAt  sql.NullTime  `db:"created_at" json:"created_at"`
	ModifiedAt sql.NullTime  `db:"modified_at" json:"modified_at"`
}

type Payment struct {
	ID         int32          `db:"id" json:"id"`
	OrderID    sql.NullInt32  `db:"order_id" json:"order_id"`
	Amount     sql.NullString `db:"amount" json:"amount"`
	Status     string         `db:"status" json:"status"`
	CreatedAt  sql.NullTime   `db:"created_at" json:"created_at"`
	ModifiedAt sql.NullTime   `db:"modified_at" json:"modified_at"`
}

type Product struct {
	ID          int32          `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Description sql.NullString `db:"description" json:"description"`
	Sku         sql.NullInt32  `db:"sku" json:"sku"`
	Category    sql.NullString `db:"category" json:"category"`
	Price       sql.NullString `db:"price" json:"price"`
	DiscountID  sql.NullInt32  `db:"discount_id" json:"discount_id"`
	CreatedAt   sql.NullTime   `db:"created_at" json:"created_at"`
	ModifiedAt  sql.NullTime   `db:"modified_at" json:"modified_at"`
}

type ShoppingSession struct {
	ID         int32          `db:"id" json:"id"`
	UserID     sql.NullInt32  `db:"user_id" json:"user_id"`
	Total      sql.NullString `db:"total" json:"total"`
	CreatedAt  sql.NullTime   `db:"created_at" json:"created_at"`
	ModifiedAt sql.NullTime   `db:"modified_at" json:"modified_at"`
}

type User struct {
	ID         int32        `db:"id" json:"id"`
	Username   string       `db:"username" json:"username"`
	Email      string       `db:"email" json:"email"`
	Password   string       `db:"password" json:"password"`
	CreatedAt  sql.NullTime `db:"created_at" json:"created_at"`
	ModifiedAt sql.NullTime `db:"modified_at" json:"modified_at"`
}
