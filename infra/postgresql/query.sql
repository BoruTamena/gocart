
-- name: GetActiveSession :one
SELECT id, user_id, total, created_at, modified_at 
FROM shopping_session
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: CreateShoppingSession :one
INSERT INTO shopping_session (user_id, total, created_at, modified_at)
VALUES ($1, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id;


-- name: AddCartItem :one
INSERT INTO cart_item (session_id, product_id, quantity, created_at, modified_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
ON CONFLICT (session_id, product_id) DO UPDATE 
SET quantity = CASE WHEN cart_item.quantity <  10  THEN cart_item.quantity + EXCLUDED.quantity
    ELSE 10 
END,
modified_at = CURRENT_TIMESTAMP 
RETURNING cart_item.quantity;


-- name: RemoveCartItem :execrows
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
);


-- name: UpdateCartItemQuantity :exec
UPDATE cart_item
SET quantity = $3, modified_at = CURRENT_TIMESTAMP
WHERE session_id = $1 AND product_id = $2
RETURNING id ;

-- name: ViewCurrentCartITem :many
SELECT * FROM  product
WHERE product.id IN (
    SELECT product_id 
    FROM cart_item 
    WHERE session_id=$1
);

-- name: IncreaseQuantity :exec
UPDATE cart_item 
SET quantity=quantity+1
WHERE session_id = $1 AND product_id = $2 ;



-- name: DecreaseQuantity :exec
UPDATE cart_item 
SET quantity=quantity-1
WHERE session_id = $1 AND product_id = $2 AND quantity > 1;




-- name: CheckoutCart :exec
BEGIN;
-- Create a new order for the user
INSERT INTO "order_details" (user_id, total, created_at)
SELECT ss.user_id, SUM(p.price * ci.quantity), CURRENT_TIMESTAMP
FROM shopping_session ss
JOIN cart_item ci ON ci.session_id = ss.id
JOIN product p ON ci.product_id = p.id
WHERE ss.user_id = $1
GROUP BY ss.user_id
RETURNING id INTO order_id;

-- Insert items from cart into order_items
INSERT INTO "order_items" (order_id, product_id, created_at)
SELECT order_id, ci.product_id, CURRENT_TIMESTAMP
FROM cart_item ci
JOIN shopping_session ss ON ci.session_id = ss.id
WHERE ss.user_id = $1;

-- Clear the user's cart
DELETE FROM cart_item 
WHERE session_id = (SELECT id FROM shopping_session WHERE user_id = $1);

COMMIT;
