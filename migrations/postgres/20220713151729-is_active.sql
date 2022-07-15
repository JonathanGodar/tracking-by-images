-- +migrate Up
ALTER TABLE trackers ADD COLUMN is_active boolean NOT NULL default false;

-- +migrate Down
ALTER TABLE trackers DROP COLUMN is_active;
