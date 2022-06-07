CREATE TYPE provider AS ENUM ('GOOGLE', 'FACEBOOK', 'APPLE');
CREATE TYPE type AS ENUM ('USER', 'ADMIN', 'DEV', 'COMPANY');

CREATE TABLE accounts (
    id char(27) PRIMARY KEY,
    email TEXT NOT NULL,
    email_verified BOOLEAN DEFAULT false,
    email_code TEXT,
    password_hash TEXT NOT NULL,
    provider provider DEFAULT NULL,
    type type NULL DEFAULT 'USER'
);

CREATE UNIQUE INDEX email_idx ON accounts (email);