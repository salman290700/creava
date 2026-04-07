-- migrate:up
CREATE TABLE refresh_token_creator (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    refresh_token VARCHAR(512) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    expired_at DATETIME NOT NULL,
    
    INDEX idx_creator_id (creator_id),
    CONSTRAINT fk_refresh_token_creator
        FOREIGN KEY (creator_id)
        REFERENCES creators(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS refresh_token_creator
