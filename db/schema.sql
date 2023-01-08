CREATE TABLE "transactions" (
  "id" int PRIMARY KEY,
  "uid" varchar NOT NULL DEFAULT (uuid_generate_v4()),
  "created" timestamptz NOT NULL DEFAULT (now()),
  "amount" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "user_id" int
);