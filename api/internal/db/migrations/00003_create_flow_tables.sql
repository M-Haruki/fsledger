-- +goose Up
CREATE TABLE flows (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  transaction_id uuid NOT NULL REFERENCES transactions (id) ON DELETE CASCADE,
  from_stock_id uuid NOT NULL REFERENCES stocks (id) ON DELETE NO ACTION,
  to_stock_id uuid NOT NULL REFERENCES stocks (id) ON DELETE NO ACTION,
  from_amount bigint NOT NULL, -- bigint = int64
  to_amount bigint NOT NULL, -- bigint = int64
  PRIMARY KEY (id)
);

CREATE TABLE flow_tags (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name text NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (name) -- TODO: For multi-user support, change to UNIQUE (user_id, name)
);

CREATE TABLE flow_tag_relations (
  flow_id uuid NOT NULL REFERENCES flows (id) ON DELETE CASCADE,
  tag_id uuid NOT NULL REFERENCES flow_tags (id) ON DELETE CASCADE,
  PRIMARY KEY (flow_id, tag_id)
);

-- +goose Down
DROP TABLE flow_tag_relations;
DROP TABLE flow_tags;
DROP TABLE flows;
