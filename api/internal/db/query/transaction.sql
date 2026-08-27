-- name: GetTransaction :one
SELECT description, occurred_at FROM transactions WHERE id = sqlc.arg(id);

-- name: CreateTransaction :one
INSERT INTO transactions (description, occurred_at) VALUES (sqlc.arg(description), sqlc.arg(occurredAt)) RETURNING id;

-- name: UpdateTransaction :execresult
UPDATE transactions SET description = sqlc.arg(description), occurred_at = sqlc.arg(occurredAt) WHERE id = sqlc.arg(id);

-- name: DeleteTransaction :execresult
DELETE FROM transactions WHERE id = sqlc.arg(id);

-- name: ListTransactionTags :many
SELECT id, name FROM transaction_tags;

-- name: GetTransactionTag :one
SELECT name FROM transaction_tags WHERE id = sqlc.arg(id);

-- name: CreateTransactionTag :one
INSERT INTO transaction_tags (name) VALUES (sqlc.arg(name)) RETURNING id;

-- name: UpdateTransactionTag :execresult
UPDATE transaction_tags SET name = sqlc.arg(name) WHERE id = sqlc.arg(id);

-- name: DeleteTransactionTag :execresult
DELETE FROM transaction_tags WHERE id = sqlc.arg(id);

-- name: CreateTransactionTagRelation :exec
INSERT INTO transaction_tag_relations (transaction_id, tag_id) VALUES (sqlc.arg(transaction_id), sqlc.arg(tag_id));

-- name: DeleteTransactionTagRelation :exec
DELETE FROM transaction_tag_relations WHERE transaction_id = sqlc.arg(transaction_id);

-- name: ListTagIDsByTransaction :many
SELECT tag_id FROM transaction_tag_relations WHERE transaction_id = sqlc.arg(id);
