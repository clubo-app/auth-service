CREATE TYPE provider AS ENUM (
    'GOOGLE', 
    'FACEBOOK', 
    'APPLE'
);

CREATE TYPE type AS ENUM (
  'USER', 
  'ADMIN', 
  'DEV', 
  'COMPANY'
);

CREATE TABLE accounts (
    id varchar(27) PRIMARY KEY,
    email TEXT NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT false,
    email_code TEXT,
    password_hash TEXT NOT NULL,
    refresh_token_generation SMALLINT NOT NULL DEFAULT 1,
    provider provider,
    type type NOT NULL DEFAULT 'USER'
);

CREATE UNIQUE INDEX email_idx ON accounts (email);
