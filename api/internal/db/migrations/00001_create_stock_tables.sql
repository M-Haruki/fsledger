-- +goose Up
CREATE TABLE stocks (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name text NOT NULL,
  has_amount boolean NOT NULL,
  currency text NOT NULL,
  description text NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (name) -- TODO: For multi-user support, change to UNIQUE (user_id, name)
);

COMMENT ON COLUMN stocks.has_amount IS 'Whether the total amount is meaningful';

CREATE TABLE stock_tags (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name text NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (name) -- TODO: For multi-user support, change to UNIQUE (user_id, name)
);

CREATE TABLE stock_tag_relations (
  stock_id uuid NOT NULL REFERENCES stocks (id) ON DELETE CASCADE,
  tag_id uuid NOT NULL REFERENCES stock_tags (id) ON DELETE CASCADE,
  PRIMARY KEY (stock_id, tag_id)
);

-- +goose Down
DROP TABLE stock_tag_relations;
DROP TABLE stock_tags;
DROP TABLE stocks;
