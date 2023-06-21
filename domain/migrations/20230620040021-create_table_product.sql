-- +migrate Up
CREATE TABLE IF NOT EXISTS product (
                         id VARCHAR(36) DEFAULT (UUID()),
                         name VARCHAR(255) NOT NULL,
                         price DECIMAL(10, 2) NOT NULL,
                         created_at timestamp NOT NULL,
                         updated_at timestamp NOT NULL,
                         deleted_at timestamp ,
                         PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS product;