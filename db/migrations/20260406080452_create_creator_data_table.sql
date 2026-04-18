-- migrate:up
CREATE TABLE creator_data (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator BIGINT NOT NULL,
    email BIGINT NOT NULL,
    phone_number BIGINT,
    image_url BIGINT,
    address BIGINT,
    version BIGINT NOT NULL DEFAULT 1,
    status BIGINT,    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_creator (creator),
    INDEX idx_email (email),
    INDEX idx_phone_number (phone_number),
    INDEX idx_address (address),
    INDEX idx_status (status),

    CONSTRAINT fk_creator_data_creator
        FOREIGN KEY (creator) REFERENCES creators(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_data_email
        FOREIGN KEY (email) REFERENCES creator_emails(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_data_phone
        FOREIGN KEY (phone_number) REFERENCES creator_contacts(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_data_address
        FOREIGN KEY (address) REFERENCES creator_addresses(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_data_status
        FOREIGN KEY (status) REFERENCES creator_statuses(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    CONSTRAINT fk_creator_data_image_url
    FOREIGN KEY (image_url) REFERENCES creator_image(id)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,

    UNIQUE KEY unique_creator_data (creator)
);

-- migrate:down
DROP TABLE IF EXISTS creator_data
