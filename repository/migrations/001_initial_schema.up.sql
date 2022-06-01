CREATE TYPE IF NOT EXISTS provider AS ENUM ('google', 'facebook', 'apple');
CREATE TYPE IF NOT EXISTS role AS ENUM ('admin', 'dev', 'user');
CREATE TYPE IF NOT EXISTS type AS ENUM ('user', 'company');

CREATE TABLE accounts (
    id char(27) PRIMARY KEY,
    email TEXT NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT false,
    email_code TEXT,
    password_hash TEXT NOT NULL,
    provider provider NULL,
    role role NULL,
    type type NULL
);

CREATE UNIQUE INDEX email_idx ON users (email);