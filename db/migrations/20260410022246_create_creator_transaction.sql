-- migrate:up
CREATE TABLE IF NOT EXISTS creator_transaction (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    balance_id BIGINT NOT NULL,
    transaction_type VARCHAR(50) NOT NULL, 
    amount DECIMAL(15,2) NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',
    reference_id BIGINT NULL,
    description VARCHAR(255),
    status VARCHAR(50) NOT NULL,
    version BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_creator_transaction_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT fk_creator_transaction_balance
        FOREIGN KEY (balance_id) REFERENCES creator_balance(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE INDEX idx_creator_transaction_creator_id 
ON creator_transaction(creator_id);

CREATE INDEX idx_creator_transaction_balance_id 
ON creator_transaction(balance_id);

CREATE INDEX idx_creator_transaction_status 
ON creator_transaction(status);

CREATE INDEX idx_creator_transaction_type 
ON creator_transaction(transaction_type);

CREATE INDEX idx_creator_transaction_created_at 
ON creator_transaction(created_at);

-- migrate:down
-- Drop indexes
DROP INDEX idx_creator_transaction_sender ON creator_transaction;
DROP INDEX idx_creator_transaction_receiver ON creator_transaction;
DROP INDEX idx_creator_transaction_balance ON creator_transaction;
DROP INDEX idx_creator_transaction_type ON creator_transaction;
DROP INDEX idx_creator_transaction_status ON creator_transaction;
DROP INDEX idx_creator_transaction_created_at ON creator_transaction;

-- Drop table
DROP TABLE IF EXISTScreator_transaction;
