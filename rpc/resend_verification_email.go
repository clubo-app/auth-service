package rpc

import (
	"context"

	ag "github.com/clubo-app/protobuf/auth"
	cg "github.com/clubo-app/protobuf/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) ResendVerificationEmail(context.Context, *ag.ResendVerificationEmailRequest) (*cg.SuccessIndicator, error) {
	return nil, status.Error(codes.Unavailable, "not yet implemented")

}
