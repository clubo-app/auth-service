package rpc

import (
	"context"

	cg "github.com/clubo-app/protobuf/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) ResendVerificationEmail(context.Context, *cg.Empty) (*cg.SuccessIndicator, error) {
	return nil, status.Error(codes.Unavailable, "not yet implemented")

}
