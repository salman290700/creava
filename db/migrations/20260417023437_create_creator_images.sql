-- migrate:up
CREATE TABLE creator_image (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_creator_id (creator_id),
    CONSTRAINT fk_creator_image_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS creator_images
