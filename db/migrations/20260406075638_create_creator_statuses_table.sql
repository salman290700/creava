-- migrate:up
CREATE TABLE creator_statuses (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    status BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_creator_id (creator_id),
    INDEX idx_status (status),

    CONSTRAINT fk_creator_status_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_status_status
        FOREIGN KEY (status) REFERENCES statuses(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS creator_statuses
