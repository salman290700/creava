-- migrate:up
CREATE TABLE commission_fee (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    transaction_id BIGINT NOT NULL,
    donation_id BIGINT NULL,
    commission_type VARCHAR(50) NOT NULL,
    commission_percentage DECIMAL(5,2) NOT NULL,
    commission_amount DECIMAL(15,2) NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',
    description VARCHAR(255),
    version BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_commission_fee_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_commission_fee_transaction
        FOREIGN KEY (transaction_id) REFERENCES creator_transaction(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_commission_fee_donation
        FOREIGN KEY (donation_id) REFERENCES creator_donation(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE INDEX idx_commission_fee_creator 
ON commission_fee(creator_id);

CREATE INDEX idx_commission_fee_transaction 
ON commission_fee(transaction_id);

CREATE INDEX idx_commission_fee_donation 
ON commission_fee(donation_id);

CREATE INDEX idx_commission_fee_created_at 
ON commission_fee(created_at);

-- migrate:down
-- Drop indexes
DROP INDEX idx_commission_fee_creator ON commission_fee;
DROP INDEX idx_commission_fee_transaction ON commission_fee;
DROP INDEX idx_commission_fee_donation ON commission_fee;
DROP INDEX idx_commission_fee_created_at ON commission_fee;

-- Drop table
DROP TABLE commission_fee;  