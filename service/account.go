package service

import (
	"context"
	"database/sql"

	"github.com/clubo-app/auth-service/dto"
	"github.com/clubo-app/auth-service/repository"
	"github.com/clubo-app/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountService interface {
	Create(context.Context, dto.Account) (repository.Account, error)
	Delete(context.Context, string) error
	Update(context.Context, dto.Account) (repository.Account, error)
	UpdateVerified(ctx context.Context, id, code string, emailVerified bool) (repository.Account, error)
	RotateEmailCode(ctx context.Context, email string) (repository.Account, error)
	EmailTaken(ctx context.Context, email string) bool
	GetById(ctx context.Context, id string) (repository.Account, error)
	GetByEmail(ctx context.Context, email string) (repository.Account, error)
}

type accountService struct {
	q *repository.Queries
}

func NewAccountService(q *repository.Queries) AccountService {
	return &accountService{q: q}
}

func (s *accountService) Create(ctx context.Context, d dto.Account) (repository.Account, error) {
	a, err := s.q.CreateAccount(ctx, repository.CreateAccountParams{
		ID:            d.ID,
		Email:         d.Email,
		EmailVerified: d.EmailVerified,
		EmailCode:     sql.NullString{String: d.EmailCode, Valid: d.EmailCode != ""},
		PasswordHash:  d.PasswordHash,
		Provider:      d.Provider,
		Role:          d.Role,
		Type:          d.Type,
	})
	if err != nil {
		return repository.Account{}, err
	}

	return a, err
}

func (s *accountService) Delete(ctx context.Context, id string) error {
	err := s.q.DeleteAccount(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *accountService) Update(ctx context.Context, d dto.Account) (repository.Account, error) {
	a, err := s.q.UpdateAccount(ctx, repository.UpdateAccountParams{
		ID:           d.ID,
		Email:        sql.NullString{String: d.Email, Valid: d.Email != ""},
		EmailCode:    sql.NullString{String: d.EmailCode, Valid: d.EmailCode != ""},
		PasswordHash: sql.NullString{String: d.PasswordHash, Valid: d.PasswordHash != ""},
		Role:         d.Role,
	})
	if err != nil {
		return repository.Account{}, err
	}

	return a, err
}

func (s *accountService) UpdateVerified(ctx context.Context, id, code string, emailVerified bool) (repository.Account, error) {
	a, err := s.q.UpdateVerified(ctx, repository.UpdateVerifiedParams{
		ID:       id,
		Verified: emailVerified,
		Code:     sql.NullString{String: code, Valid: true},
	})
	if err != nil {
		return repository.Account{}, err
	}

	return a, err
}

func (s *accountService) RotateEmailCode(ctx context.Context, email string) (repository.Account, error) {
	code, err := utils.GenerateOTP(4)
	if err != nil {
		return repository.Account{}, status.Error(codes.Internal, "Failed to generate Email Code")
	}

	a, err := s.q.UpdateEmailCode(ctx, repository.UpdateEmailCodeParams{
		EmailCode: sql.NullString{String: code, Valid: true},
		Email:     email,
	})
	if err != nil {
		return repository.Account{}, err
	}

	return a, err
}

func (s *accountService) EmailTaken(ctx context.Context, email string) bool {
	t, err := s.q.EmailTaken(ctx, email)
	if err != nil {
		return false
	}
	return t
}

func (s *accountService) GetById(ctx context.Context, id string) (repository.Account, error) {
	a, err := s.q.GetAccount(ctx, id)
	if err != nil {
		return repository.Account{}, err
	}

	return a, err
}

func (s *accountService) GetByEmail(ctx context.Context, email string) (repository.Account, error) {
	a, err := s.q.GetAccountByEmail(ctx, email)
	if err != nil {
		return repository.Account{}, err
	}

	return a, err
}
