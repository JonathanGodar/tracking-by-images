
-- +migrate Up
ALTER TABLE users RENAME COLUMN "username" TO "email";

-- +migrate Down
ALTER TABLE users RENAME COLUMN "email" TO "username";
