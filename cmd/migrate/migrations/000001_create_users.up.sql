CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users(
  id BIGSERIAL PRIMARY KEY,
  email citext UNIQUE NOT NULL,
  username VARCHAR(255) NOT NULL,
  password bytea NOT NULL,
  created_ad timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
