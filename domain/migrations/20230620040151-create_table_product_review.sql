-- +migrate Up
CREATE TABLE IF NOT EXISTS product_review (
                                  id VARCHAR(36) DEFAULT (UUID()),
                                  product_id VARCHAR(36) NOT NULL,
                                  member_id VARCHAR(36) NOT NULL,
                                  desc_review TEXT,
                                  created_at timestamp NOT NULL,
                                  updated_at timestamp NOT NULL,
                                  deleted_at timestamp NULL,
                                  PRIMARY KEY (id),
                                  FOREIGN KEY (product_id) REFERENCES product(id),
                                  FOREIGN KEY (member_id) REFERENCES member(id)
  );

-- +migrate Down
DROP TABLE IF EXISTS product_review;