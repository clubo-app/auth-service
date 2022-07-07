package rpc

import (
	"context"
	"net/mail"

	"github.com/clubo-app/packages/utils"
	ag "github.com/clubo-app/protobuf/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) LoginUser(ctx context.Context, req *ag.LoginUserRequest) (*ag.LoginUserResponse, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}
	a, err := s.ac.GetByEmail(ctx, req.Email)

	pwEqual := s.pw.CheckPasswordHash(req.Password, a.PasswordHash)
	if !pwEqual {
		return nil, status.Error(codes.InvalidArgument, "Invalid Password")
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
