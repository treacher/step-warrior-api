
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_plans (
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL REFERENCES users(id),
  plan_id bigint NOT NULL REFERENCES plans(id),
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE INDEX user_plans_user_id_idx ON user_plans (user_id);
CREATE UNIQUE INDEX user_plans_user_id_plan_id_idx ON user_plans (user_id, plan_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_plans;
