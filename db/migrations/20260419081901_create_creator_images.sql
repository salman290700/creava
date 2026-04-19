-- migrate:up
CREATE TABLE IF NOT EXISTS creator_image (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    image_url BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_creator_id (creator_id),
    CONSTRAINT fk_image_url
        FOREIGN KEY (image_url) REFERENCES images(id)
        ON DELETE RESTRICT,
    CONSTRAINT fk_creator_image_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
)

-- migrate:down
DROP TABLE IF EXISTS creator_image
