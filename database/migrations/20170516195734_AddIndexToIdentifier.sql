
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE UNIQUE INDEX users_identifier ON users (identifier);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX users_identifier;
