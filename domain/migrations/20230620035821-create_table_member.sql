-- +migrate Up
CREATE TABLE IF NOT EXISTS member(
                        id VARCHAR(36) DEFAULT (UUID()),
                        username VARCHAR(255) NOT NULL,
                        gender VARCHAR(255) NOT NULL,
                        skintype VARCHAR(255) NOT NULL,
                        skincolor VARCHAR(255) NOT NULL,
                        created_at timestamp NOT NULL,
                        updated_at timestamp NOT NULL,
                        deleted_at timestamp ,
                        PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS member;