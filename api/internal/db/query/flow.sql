-- name: ListFlowByTransaction :many
SELECT id, from_stock_id, to_stock_id, from_amount, to_amount FROM flows WHERE transaction_id = sqlc.arg(id);

-- name: CreateFlow :one
INSERT INTO flows (transaction_id, from_stock_id, to_stock_id, from_amount, to_amount) VALUES (sqlc.arg(transactionId), sqlc.arg(fromStockId), sqlc.arg(toStockId), sqlc.arg(from_amount), sqlc.arg(to_amount)) RETURNING id;

-- name: DeleteFlowByTransaction :exec
DELETE FROM flows WHERE transaction_id = sqlc.arg(id);

-- name: ListFlowTags :many
SELECT id, name FROM flow_tags;

-- name: GetFlowTag :one
SELECT name FROM flow_tags WHERE id = sqlc.arg(id);

-- name: CreateFlowTag :one
INSERT INTO flow_tags (name) VALUES (sqlc.arg(name)) RETURNING id;

-- name: UpdateFlowTag :execresult
UPDATE flow_tags SET name = sqlc.arg(name) WHERE id = sqlc.arg(id);

-- name: DeleteFlowTag :execresult
DELETE FROM flow_tags WHERE id = sqlc.arg(id);

-- name: CreateFlowTagRelation :exec
INSERT INTO flow_tag_relations (flow_id, tag_id) VALUES (sqlc.arg(flow_id), sqlc.arg(tag_id));

-- name: DeleteFlowTagRelation :exec
DELETE FROM flow_tag_relations WHERE flow_id = sqlc.arg(flow_id);

-- name: ListTagIDsByFlow :many
SELECT tag_id FROM flow_tag_relations WHERE flow_id = sqlc.arg(id);
