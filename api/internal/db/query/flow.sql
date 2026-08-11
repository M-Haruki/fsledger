-- name: ListFlowByTransaction :many
SELECT id, from_stock_id, to_stock_id, amount FROM flows WHERE transaction_id = sqlc.arg(id);

-- name: CreateFlow :one
INSERT INTO flows (transaction_id, from_stock_id, to_stock_id, amount) VALUES (sqlc.arg(transaction_id), sqlc.arg(from_stock_id), sqlc.arg(to_stock_id), sqlc.arg(amount)) RETURNING id;

-- name: UpdateFlow :exec
UPDATE flows SET from_stock_id = sqlc.arg(from_stock_id), to_stock_id = sqlc.arg(to_stock_id), amount = sqlc.arg(amount) WHERE id = sqlc.arg(id);

-- name: DeleteFlow :exec
DELETE FROM flows WHERE id = sqlc.arg(id);

-- name: ListFlowTags :many
SELECT id, name FROM flow_tags;

-- name: GetFlowTag :one
SELECT id, name FROM flow_tags WHERE id = sqlc.arg(id);

-- name: CreateFlowTag :one
INSERT INTO flow_tags (name) VALUES (sqlc.arg(name)) RETURNING id;

-- name: UpdateFlowTag :exec
UPDATE flow_tags SET name = sqlc.arg(name) WHERE id = sqlc.arg(id);

-- name: DeleteFlowTag :exec
DELETE FROM flow_tags WHERE id = sqlc.arg(id);

-- name: ListTagsByFlow :many
SELECT r.tag_id AS id, t.name AS name FROM flow_tag_relations r JOIN flow_tags t ON r.tag_id = t.id WHERE r.flow_id = sqlc.arg(id);
