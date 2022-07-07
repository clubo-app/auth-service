package repository

import (
	"context"
	"fmt"
	"log"
	"net/url"

	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leporo/sqlf"
)

const (
	TABLE_NAME = "accounts"
)

type AccountRepository struct {
	pool    *pgxpool.Pool
	querier Querier
}

func NewAccountRepository(dbUser, dbPW, dbName, dbHost string, dbPort uint16) (*AccountRepository, error) {
	urlStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPW, dbHost, fmt.Sprint(dbPort), dbName)
	pgURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	connURL := *pgURL
	if connURL.Scheme == "cockroachdb" {
		connURL.Scheme = "postgres"
	}

	c, err := pgxpool.ParseConfig(connURL.String())
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("pgx connection error: %w", err)
	}

	err = migrateSchema(connURL)
	if err != nil {
		log.Printf("Schema validation error: %v", err)
	}

	return &AccountRepository{
		pool:    pool,
		querier: New(pool),
	}, nil
}

func (d AccountRepository) Close() {
	d.pool.Close()
}

const columns = "id, email, email_verified, email_code, password_hash, provider, type"

func (r AccountRepository) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	return r.querier.CreateAccount(ctx, arg)
}

type UpdateAccountParams struct {
	Email        string
	EmailCode    string
	PasswordHash string
	ID           string
}

func (r AccountRepository) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.Update(TABLE_NAME)

	if arg.Email != "" {
		b = b.Set("email", arg.Email)
	}
	if arg.EmailCode != "" {
		b = b.Set("email_code", arg.EmailCode)
	}
	if arg.PasswordHash != "" {
		b = b.Set("password_hash", arg.PasswordHash)
	}

	b.
		Where("id = ?", arg.ID).
		Returning(columns)

	row := r.pool.QueryRow(ctx, b.String(), b.Args()...)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.EmailVerified,
		&i.EmailCode,
		&i.PasswordHash,
		&i.Provider,
		&i.Type,
	)
	return i, err

}

func (r AccountRepository) UpdateEmailCode(ctx context.Context, arg UpdateEmailCodeParams) (Account, error) {
	return r.querier.UpdateEmailCode(ctx, arg)
}

func (r AccountRepository) UpdateVerified(ctx context.Context, arg UpdateVerifiedParams) (Account, error) {
	return r.querier.UpdateVerified(ctx, arg)
}

func (r AccountRepository) DeleteAccount(ctx context.Context, id string) error {
	return r.querier.DeleteAccount(ctx, id)
}

func (r AccountRepository) EmailTaken(ctx context.Context, email string) (bool, error) {
	return r.querier.EmailTaken(ctx, email)
}

func (r AccountRepository) GetAccount(ctx context.Context, id string) (Account, error) {
	return r.querier.GetAccount(ctx, id)
}

func (r AccountRepository) GetAccountByEmail(ctx context.Context, email string) (Account, error) {
	return r.querier.GetAccountByEmail(ctx, email)
}
