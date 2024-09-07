
-- name: GetActiveSession :one
SELECT id, user_id, total, created_at, modified_at 
FROM shopping_session
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: CreateShoppingSession :exec
INSERT INTO shopping_session (user_id, total, created_at, modified_at)
VALUES ($1, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id;


-- name: AddCartItem :exec
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
) ,
updated_session AS (
    UPDATE  shopping_session 
    SET total= (
        SELECT COALESCE(SUM(quantity),0)
        FROM cart_item
        WHERE session_id=$2
    )
    WHERE id=$2 RETURNING id,total
)
UPDATE shopping_session 
SET total=0 
WHERE id =$2 AND NOT EXISTS  (
   SELECT 1 FROM cart_item WHERE session_id = $2
);


-- name: UpdateCartItemQuantity :exec
UPDATE cart_item
SET quantity = $3, modified_at = CURRENT_TIMESTAMP
WHERE session_id = $1 AND product_id = $2
RETURNING id ;

-- name: ViewCurrentCartITem :exec 
SELECT * FROM  product
WHERE product.id IN (
    SELECT product_id 
    FROM cart_item 
    WHERE session_id=$1
);


