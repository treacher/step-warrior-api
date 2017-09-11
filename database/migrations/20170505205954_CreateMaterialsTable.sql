
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE materials (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  rarity varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE materials;
