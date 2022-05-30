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
		Provider:      repository.ProviderGoogle,
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

	return &ag.LoginUserResponse{
		Token:   t,
		Account: a.ToGRPCAccount(),
	}, nil

}
