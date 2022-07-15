
-- +migrate Up
ALTER TABLE users ADD CONSTRAINT email_unique UNIQUE (email);

-- +migrate Down
ALTER TABLE users DROP CONSTRAINT email_unique;
