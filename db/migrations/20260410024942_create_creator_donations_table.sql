-- migrate:up
CREATE TABLE creator_donation (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    sender_creator_id BIGINT NULL,
    receiver_creator_id BIGINT NOT NULL,
    transaction_id BIGINT NOT NULL,
    amount DECIMAL(15,2) NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',  
    donor_name VARCHAR(100),
    donor_email VARCHAR(150),
    donor_message TEXT,
    message VARCHAR(255),
    payment_method VARCHAR(50),
    external_reference VARCHAR(100),
    status VARCHAR(50) NOT NULL,
    version BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_creator_donation_sender
        FOREIGN KEY (sender_creator_id) REFERENCES creators(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_donation_receiver
        FOREIGN KEY (receiver_creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_donation_transaction
        FOREIGN KEY (transaction_id) REFERENCES creator_transaction(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE INDEX idx_creator_donation_sender
ON creator_donation(sender_creator_id);

CREATE INDEX idx_creator_donation_receiver
ON creator_donation(receiver_creator_id);

CREATE INDEX idx_creator_donation_transaction
ON creator_donation(transaction_id);

CREATE INDEX idx_creator_donation_status
ON creator_donation(status);

CREATE INDEX idx_creator_donation_created_at
ON creator_donation(created_at);

-- migrate:down
-- Drop indexes
DROP INDEX idx_creator_donation_sender ON creator_donation;
DROP INDEX idx_creator_donation_receiver ON creator_donation;
DROP INDEX idx_creator_donation_transaction ON creator_donation;
DROP INDEX idx_creator_donation_status ON creator_donation;
DROP INDEX idx_creator_donation_created_at ON creator_donation;

-- Drop table
DROP TABLE creator_donation;
