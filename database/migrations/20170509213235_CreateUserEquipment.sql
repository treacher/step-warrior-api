
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_equipment (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES users(id),
  equipment_id bigint NOT NULL REFERENCES equipment(id),
  attributes json NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE INDEX user_equipment_user_id_idx ON user_equipment (user_id);
CREATE INDEX user_equipment_user_id_equipment_id_idx ON user_equipment (user_id, equipment_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_equipment;
