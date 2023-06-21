-- +migrate Up
CREATE TABLE IF NOT EXISTS like_review (
    id VARCHAR(36) DEFAULT (UUID()),
    product_review_id VARCHAR(36) NOT NULL,
    member_id VARCHAR(36) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    deleted_at timestamp NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (product_review_id) REFERENCES product_review(id),
    FOREIGN KEY (member_id) REFERENCES member(id)
    );

-- +migrate Down
DROP TABLE IF EXISTS like_review;