
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE plans (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  rarity varchar(255) NOT NULL,
  equipment_id bigint NOT NULL references equipment(id),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE UNIQUE INDEX plans_equipment_id_idx ON plans (equipment_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE plans;
