-- migrate:up
CREATE TABLE contacts (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    phone_number VARCHAR(20) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX phone_number_contacts (phone_number)
);

-- migrate:down
DROP TABLE IF EXISTS contacts