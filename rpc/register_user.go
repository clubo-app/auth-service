package rpc

import (
	"context"
	"net/mail"

	"github.com/clubo-app/auth-service/dto"
	"github.com/clubo-app/auth-service/repository"
	"github.com/clubo-app/packages/utils"
	ag "github.com/clubo-app/protobuf/auth"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) RegisterUser(ctx context.Context, req *ag.RegisterUserRequest) (*ag.LoginUserResponse, error) {
	hash, err := s.pw.HashPassword(req.Password)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	code, _ := utils.GenerateOTP(4)

	da := dto.Account{
		ID:            ksuid.New().String(),
		Email:         req.Email,
		EmailVerified: false,
		EmailCode:     code,
		PasswordHash:  hash,
		Provider:      repository.NullProvider{Valid: false},
		Type:          repository.TypeUSER,
	}

	a, err := s.ac.Create(ctx, da)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	at, err := s.token.NewAccessToken(a)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	rt, err := s.token.NewRefreshToken(a.ID, a.RefreshTokenGeneration)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &ag.LoginUserResponse{
		Tokens: &ag.TokenResponse{
			AccessToken:  at,
			RefreshToken: rt,
		},
		Account: a.ToGRPCAccount(),
	}, nil
}
