-- migrate:up
CREATE TABLE creator_contacts (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    contact_id BIGINT NOT NULL,
    creator_id BIGINT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_contact_id (contact_id),
    INDEX idx_creator_id (creator_id),

    CONSTRAINT fk_creator_contacts_contact
        FOREIGN KEY (contact_id) REFERENCES contacts(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_creator_contacts_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
);

-- migrate:down
DROP TABLE IF EXISTS creator_contacts
