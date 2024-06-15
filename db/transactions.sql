CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    wallet_id INT REFERENCES wallets(id),
    type VARCHAR(10) CHECK (type IN ('top-up', 'transfer')),
    amount NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

