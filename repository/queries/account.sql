-- name: CreateAccount :one
INSERT INTO accounts (
    id,
    email,
    email_verified,
    email_code,
    password_hash,
    provider,
    type
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByEmail :one
SELECT * FROM accounts
WHERE email = $1 LIMIT 1;

-- name: EmailTaken :one
select exists(select 1 from accounts where email=$1) AS "exists";

-- name: UpdateEmailCode :one
UPDATE accounts 
SET email_code = $1
WHERE email = $2
RETURNING *;

-- name: UpdateVerified :one
UPDATE accounts 
SET email_verified = IF(email_code = sqlc.arg('code'), sqlc.arg('verified')::boolean, email_code)
WHERE id = sqlc.arg('id')
RETURNING *;

