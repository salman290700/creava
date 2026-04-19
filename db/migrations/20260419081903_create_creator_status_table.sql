-- migrate:up
CREATE TABLE creator_status (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    status_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_creator_status_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_creator_status_status
        FOREIGN KEY (status_id) REFERENCES statuses(id)
        ON DELETE RESTRICT
);

-- migrate:down

DROP TABLES IF EXISTS creator_status