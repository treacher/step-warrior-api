
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_materials (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES users(id),
  material_id bigint NOT NULL REFERENCES materials(id),
  quantity bigint NOT NULL DEFAULT 1,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE INDEX user_materials_user_id_idx ON user_materials (user_id);
CREATE UNIQUE INDEX user_materials_user_id_material_id_idx ON user_materials (user_id, material_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_materials;
