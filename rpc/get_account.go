package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	ag "github.com/clubo-app/protobuf/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) GetAccount(ctx context.Context, req *ag.GetAccountRequest) (*ag.Account, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	a, err := s.ac.GetById(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return a.ToGRPCAccount(), nil
}
