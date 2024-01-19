CREATE TABLE IF NOT EXISTS submissions(
    id BIGSERIAL PRIMARY KEY,
    relay_id INTEGER NOT NULL,
    bid_block_hash BYTEA NOT NULL,
    is_delivered BOOLEAN NOT NULL,
    is_optimistic BOOLEAN NOT NULL,
    submitted_at TIMESTAMP NOT NULL
);
