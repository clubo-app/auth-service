// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: account.sql

package repository

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    id,
    email,
    email_verified,
    email_code,
    password_hash,
    provider,
    role,
    type
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, email, email_verified, email_code, password_hash, provider, role, type
`

type CreateAccountParams struct {
	ID            string
	Email         string
	EmailVerified bool
	EmailCode     sql.NullString
	PasswordHash  string
	Provider      Provider
	Role          Role
	Type          Type
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount,
		arg.ID,
		arg.Email,
		arg.EmailVerified,
		arg.EmailCode,
		arg.PasswordHash,
		arg.Provider,
		arg.Role,
		arg.Type,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.EmailCode,
		&i.PasswordHash,
		&i.Provider,
		&i.Role,
		&i.Type,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteAccount, id)
	return err
}

const emailTaken = `-- name: EmailTaken :one
select exists(select 1 from accounts where email=$1) AS "exists"
`

func (q *Queries) EmailTaken(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, emailTaken, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, email, email_verified, email_code, password_hash, provider, role, type FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id string) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.EmailCode,
		&i.PasswordHash,
		&i.Provider,
		&i.Role,
		&i.Type,
	)
	return i, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT id, email, email_verified, email_code, password_hash, provider, role, type FROM accounts
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (Account, error) {
	row := q.db.QueryRow(ctx, getAccountByEmail, email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.EmailCode,
		&i.PasswordHash,
		&i.Provider,
		&i.Role,
		&i.Type,
	)
	return i, err
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts SET
    email = $1,
    email_code = $2,
    password_hash = $3,
    role = $4
WHERE id = $5
RETURNING id, email, email_verified, email_code, password_hash, provider, role, type
`

type UpdateAccountParams struct {
	Email        sql.NullString
	EmailCode    sql.NullString
	PasswordHash sql.NullString
	Role         Role
	ID           string
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount,
		arg.Email,
		arg.EmailCode,
		arg.PasswordHash,
		arg.Role,
		arg.ID,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.EmailCode,
		&i.PasswordHash,
		&i.Provider,
		&i.Role,
		&i.Type,
	)
	return i, err
}

const updateEmailCode = `-- name: UpdateEmailCode :one
UPDATE accounts 
SET email_code = $1
WHERE email = $2
RETURNING id, email, email_verified, email_code, password_hash, provider, role, type
`

type UpdateEmailCodeParams struct {
	EmailCode sql.NullString
	Email     string
}

func (q *Queries) UpdateEmailCode(ctx context.Context, arg UpdateEmailCodeParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateEmailCode, arg.EmailCode, arg.Email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.EmailCode,
		&i.PasswordHash,
		&i.Provider,
		&i.Role,
		&i.Type,
	)
	return i, err
}

const updateVerified = `-- name: UpdateVerified :one
UPDATE accounts 
SET email_verified = IF(email_code = $1, $2::boolean, email_code)
WHERE id = $3
RETURNING id, email, email_verified, email_code, password_hash, provider, role, type
`

type UpdateVerifiedParams struct {
	Code     sql.NullString
	Verified bool
	ID       string
}

func (q *Queries) UpdateVerified(ctx context.Context, arg UpdateVerifiedParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateVerified, arg.Code, arg.Verified, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.EmailCode,
		&i.PasswordHash,
		&i.Provider,
		&i.Role,
		&i.Type,
	)
	return i, err
}
