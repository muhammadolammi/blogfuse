-- +goose Up
-- Add "API KEY" column to the "users" table
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) NOT NULL UNIQUE DEFAULT encode(sha256(random()::text::bytea), 'hex');

-- +goose Down
-- Remove "API KEY" column from the "users" table
ALTER TABLE users
DROP COLUMN api_key;
