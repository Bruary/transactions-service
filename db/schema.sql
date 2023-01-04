CREATE TABLE transactions (
  id   BIGSERIAL PRIMARY KEY,
  uid text NOT NULL default gen_random_uuid(),
  created timestamp NOT NULL default NOW(),
  amount text NOT NULL,
  currency text NOT NULL
);
