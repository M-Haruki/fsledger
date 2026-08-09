-- +goose Up
CREATE TABLE transactions (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  description text NOT NULL,
  occurred_at timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (id)
);

CREATE TABLE transaction_tags (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name text NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (name) -- TODO: For multi-user support, change to UNIQUE (user_id, name)
);

CREATE TABLE transaction_tag_relations (
  transaction_id uuid NOT NULL REFERENCES transactions (id) ON DELETE CASCADE,
  tag_id uuid NOT NULL REFERENCES transaction_tags (id) ON DELETE CASCADE,
  PRIMARY KEY (transaction_id, tag_id)
);

-- +goose Down
DROP TABLE transaction_tag_relations;
DROP TABLE transaction_tags;
DROP TABLE transactions;
