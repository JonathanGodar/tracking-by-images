
-- +migrate Up
ALTER TABLE trackers ADD COLUMN url varchar NOT NULL;

-- +migrate Down
ALTER TABLE trackers DROP COLUMN url;
