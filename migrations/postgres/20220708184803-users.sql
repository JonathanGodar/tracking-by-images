
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
				ID uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
				username varchar(255) NOT NULL,
				password_hash varchar(60) NOT NULL
);

CREATE TABLE trackers (
				ID uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
				times_accessed INT NOT NULL,
				owner_id uuid NOT NULL, 
				CONSTRAINT fk_owner
								FOREIGN KEY (owner_id)
								REFERENCES users(id)
);

-- +migrate Down
DROP TABLE trackers;
DROP TABLE users;
