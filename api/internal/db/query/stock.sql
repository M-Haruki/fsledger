-- name: ListStocks :many
SELECT id, name, has_amount, currency, currency_exponent FROM stocks;

-- name: GetStock :one
SELECT name, has_amount, currency, currency_exponent, description FROM stocks WHERE id = sqlc.arg(id);

-- name: CreateStock :one
INSERT INTO stocks (name, has_amount, currency, currency_exponent, description) VALUES (sqlc.arg(name), sqlc.arg(hasAmount), sqlc.arg(currency), sqlc.arg(currency_exponent), sqlc.arg(description)) RETURNING id;

-- name: UpdateStock :execresult
UPDATE stocks SET name = sqlc.arg(name), has_amount = sqlc.arg(hasAmount), currency = sqlc.arg(currency), currency_exponent = sqlc.arg(currency_exponent), description = sqlc.arg(description) WHERE id = sqlc.arg(id);

-- name: DeleteStock :execresult
DELETE FROM stocks WHERE id = sqlc.arg(id);

-- name: ListStockTags :many
SELECT id, name FROM stock_tags;

-- name: GetStockTag :one
SELECT name FROM stock_tags WHERE id = sqlc.arg(id);

-- name: CreateStockTag :one
INSERT INTO stock_tags (name) VALUES (sqlc.arg(name)) RETURNING id;

-- name: UpdateStockTag :execresult
UPDATE stock_tags SET name = sqlc.arg(name) WHERE id = sqlc.arg(id);

-- name: DeleteStockTag :execresult
DELETE FROM stock_tags WHERE id = sqlc.arg(id);

-- name: CreateStockTagRelation :exec
INSERT INTO stock_tag_relations (stock_id, tag_id) VALUES (sqlc.arg(stock_id), sqlc.arg(tag_id));

-- name: DeleteStockTagRelation :exec
DELETE FROM stock_tag_relations WHERE stock_id = sqlc.arg(stock_id);

-- name: ListTagIDsByStock :many
SELECT tag_id FROM stock_tag_relations WHERE stock_id = sqlc.arg(id);
