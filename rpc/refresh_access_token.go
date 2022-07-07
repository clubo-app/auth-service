package rpc

import (
	"context"
	"log"

	"github.com/clubo-app/packages/utils"
	ag "github.com/clubo-app/protobuf/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) RefreshAccessToken(ctx context.Context, req *ag.RefreshAccessTokenRequest) (*ag.TokenResponse, error) {
	p, err := s.token.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	a, err := s.ac.GetById(ctx, p.Subject)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	log.Println(a.RefreshTokenGeneration)
	log.Println(p.Generation)

	if a.RefreshTokenGeneration != p.Generation {
		return nil, status.Error(codes.InvalidArgument, "Invalid Token Generation Family")
	}

	t, err := s.token.NewAccessToken(a)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &ag.TokenResponse{
		AccessToken:  t,
		RefreshToken: req.RefreshToken,
	}, nil
}
