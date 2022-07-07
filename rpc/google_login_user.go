package rpc

import (
	"context"

	"github.com/clubo-app/auth-service/dto"
	"github.com/clubo-app/auth-service/repository"
	"github.com/clubo-app/packages/utils"
	ag "github.com/clubo-app/protobuf/auth"
)

func (s *authServer) GoogleLoginUser(ctx context.Context, req *ag.GoogleLoginUserRequest) (*ag.LoginUserResponse, error) {
	claims, err := s.goog.ValidateGoogleJWT(req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	code, _ := utils.GenerateOTP(4)

	da := dto.Account{
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
		EmailCode:     code,
		Provider:      repository.NullProvider{Valid: true, Provider: repository.ProviderGOOGLE},
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
