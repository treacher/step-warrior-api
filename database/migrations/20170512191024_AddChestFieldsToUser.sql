
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users ADD chest_count bigint;
ALTER TABLE users ADD last_fetched_chests_at timestamp;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE users DROP chest_count;
ALTER TABLE users DROP last_fetched_chests_at;
