CREATE TABLE IF NOT EXISTS relays(
    id SERIAL PRIMARY KEY,
    url TEXT UNIQUE NOT NULL
);

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

CREATE TABLE IF NOT EXISTS submissions(
  id BIGSERIAL PRIMARY KEY,
  relay_id INTEGER NOT NULL,
  bid_block_hash BYTEA NOT NULL,
  is_delivered BOOLEAN NOT NULL,
  is_optimistic BOOLEAN NOT NULL,
  submitted_at TIMESTAMP NOT NULL
);
