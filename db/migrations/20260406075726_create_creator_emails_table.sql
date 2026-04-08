-- migrate:up
CREATE TABLE creator_emails (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    email BIGINT NOT NULL,
    creator_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_email (email),
    INDEX idx_creator_id (creator_id),

    CONSTRAINT fk_creator_emails_email
        FOREIGN KEY (email) REFERENCES emails(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_emails_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    UNIQUE KEY unique_creator_email (creator_id, email)
);

-- migrate:down
DROP TABLE IF EXISTS creator_emails
