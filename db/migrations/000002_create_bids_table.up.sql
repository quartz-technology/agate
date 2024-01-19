CREATE TABLE IF NOT EXISTS bids(
    block_hash BYTEA PRIMARY KEY,
    slot BIGINT NOT NULL,
    parent_hash BYTEA NOT NULL,
    fee_recipient BYTEA NOT NULL,
    gas_limit BIGINT NOT NULL,
    gas_used BIGINT NOT NULL,
    value BYTEA NOT NULL,
    num_tx INT NOT NULL,
    proposer BYTEA NOT NULL,
    builder BYTEA NOT NULL
);
