-- name: GetTransaction :one
SELECT description, occurred_at FROM transactions WHERE id = sqlc.arg(id);

-- name: CreateTransaction :one
INSERT INTO transactions (description, occurred_at) VALUES (sqlc.arg(description), sqlc.arg(occurred_at)) RETURNING id;

-- name: UpdateTransaction :exec
UPDATE transactions SET description = sqlc.arg(description), occurred_at = sqlc.arg(occurred_at) WHERE id = sqlc.arg(id);

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE id = sqlc.arg(id);

-- name: ListTransactionTags :many
SELECT id, name FROM transaction_tags;

-- name: GetTransactionTag :one
SELECT id, name FROM transaction_tags WHERE id = sqlc.arg(id);

-- name: CreateTransactionTag :one
INSERT INTO transaction_tags (name) VALUES (sqlc.arg(name)) RETURNING id;

-- name: UpdateTransactionTag :exec
UPDATE transaction_tags SET name = sqlc.arg(name) WHERE id = sqlc.arg(id);

-- name: DeleteTransactionTag :exec
DELETE FROM transaction_tags WHERE id = sqlc.arg(id);

-- name: ListTagsByTransaction :many
SELECT r.tag_id AS id, t.name AS name FROM transaction_tag_relations r JOIN transaction_tags t ON r.tag_id = t.id WHERE r.transaction_id = sqlc.arg(id);
