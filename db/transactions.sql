CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    wallet_id INT REFERENCES wallets(id),
    type VARCHAR(10) CHECK (type IN ('top-up', 'transfer')),
    amount NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    wallet_id INT REFERENCES wallets(id), -- Wallet yang terkait dengan transaksi
    type VARCHAR(10) CHECK (type IN ('top-up', 'transfer', 'receive')), -- Jenis transaksi: top-up, transfer, atau menerima
    amount NUMERIC(15, 2) NOT NULL, -- Jumlah transaksi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Waktu transaksi
    sender_wallet_id INT REFERENCES wallets(id), -- Wallet pengirim dalam transaksi transfer
    receiver_wallet_id INT REFERENCES wallets(id) -- Wallet penerima dalam transaksi transfer
);
