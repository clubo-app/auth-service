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

func (s *authServer) RegisterUser(ctx context.Context, req *ag.RegisterUserRequest) (*ag.RegisterUserResponse, error) {
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
		Role:          repository.RoleUser,
		Type:          repository.TypeUser,
	}

	a, err := s.ac.Create(ctx, da)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	t, err := s.token.NewJWT(a)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &ag.RegisterUserResponse{
		Token:   t,
		Account: a.ToGRPCAccount(),
	}, nil
}
