CREATE TYPE provider AS ENUM ('GOOGLE', 'FACEBOOK', 'APPLE');
CREATE TYPE role AS ENUM ('ADMIN', 'DEV', 'USER');
CREATE TYPE type AS ENUM ('USER', 'COMPANY');

CREATE TABLE accounts (
    id char(27) PRIMARY KEY,
    email TEXT NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT false,
    email_code TEXT,
    password_hash TEXT NOT NULL,
    provider provider DEFAULT NULL,
    role role NULL  DEFAULT "USER",
    type type NULL  DEFAULT "USER"
);

CREATE UNIQUE INDEX email_idx ON accounts (email);