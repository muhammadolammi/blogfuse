-- +goose Up
-- add last fetched column to feeds table

ALTER TABLE feeds
ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE feeds
DROP COLUMN last_fetched_at ;