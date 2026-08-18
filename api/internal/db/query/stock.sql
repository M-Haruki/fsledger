-- name: ListStocks :many
SELECT id, name FROM stocks;

-- name: GetStock :one
SELECT id, name, has_amount, currency, description FROM stocks WHERE id = sqlc.arg(id);

-- name: CreateStock :one
INSERT INTO stocks (name, has_amount, currency, description) VALUES (sqlc.arg(name), sqlc.arg(hasAmount), sqlc.arg(currency), sqlc.arg(description)) RETURNING id;

-- name: UpdateStock :exec
UPDATE stocks SET name = sqlc.arg(name), has_amount = sqlc.arg(hasAmount), currency = sqlc.arg(currency), description = sqlc.arg(description) WHERE id = sqlc.arg(id);

-- name: DeleteStock :execresult
DELETE FROM stocks WHERE id = sqlc.arg(id);

-- name: ListStockTags :many
SELECT id, name FROM stock_tags;

-- name: GetStockTag :one
SELECT id, name FROM stock_tags WHERE id = sqlc.arg(id);

-- name: CreateStockTag :one
INSERT INTO stock_tags (name) VALUES (sqlc.arg(name)) RETURNING id;

-- name: UpdateStockTag :exec
UPDATE stock_tags SET name = sqlc.arg(name) WHERE id = sqlc.arg(id);

-- name: DeleteStockTag :exec
DELETE FROM stock_tags WHERE id = sqlc.arg(id);

-- name: CreateStockTagRelation :exec
INSERT INTO stock_tag_relations (stock_id, tag_id) VALUES (sqlc.arg(stock_id), sqlc.arg(tag_id));

-- name: DeleteStockTagRelation :exec
DELETE FROM stock_tag_relations WHERE stock_id = sqlc.arg(stock_id) AND tag_id = sqlc.arg(tag_id);

-- name: ListTagsByStock :many
SELECT r.tag_id AS id, t.name AS name FROM stock_tag_relations r JOIN stock_tags t ON r.tag_id = t.id WHERE r.stock_id = sqlc.arg(id);
