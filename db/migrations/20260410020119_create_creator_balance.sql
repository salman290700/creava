-- migrate:up
CREATE TABLE creator_balance (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    creator_id BIGINT NOT NULL,
    balance DECIMAL(15,2) NOT NULL DEFAULT 0,
    pending_balance DECIMAL(15,2) NOT NULL DEFAULT 0,
    withdrawn_balance DECIMAL(15,2) NOT NULL DEFAULT 0,
    currency VARCHAR(10) NOT NULL DEFAULT 'IDR',
    version BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_creator_balance_creator
        FOREIGN KEY (creator_id) REFERENCES creators(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,

    CONSTRAINT unique_creator_balance UNIQUE (creator_id)
);

CREATE INDEX idx_creator_balance_creator_id 
ON creator_balance(creator_id);

-- migrate:down
-- Drop index
DROP INDEX idx_creator_balance_creator_id ON creator_balance;

-- Drop table
DROP TABLE creator_balance;

