-- name: CreateProduct :one
INSERT INTO products
(id, name, description, price, category, image_url, created_at, updated_at)
 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
 RETURNING *;

-- name: ListProducts :many
Select * from products
WHERE name = $1 AND description = $2 AND category = $3 AND price = $4
ORDER BY id;

-- name: GetProduct :one
Select * from products
WHERE name = $1 AND description = $2 AND category = $3 AND price = $4
LIMIT 1;