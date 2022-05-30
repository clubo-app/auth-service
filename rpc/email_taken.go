package rpc

import (
	"context"

	ag "github.com/clubo-app/protobuf/auth"
)

func (s *authServer) EmailTaken(ctx context.Context, req *ag.EmailTakenRequest) (*ag.EmailTakenResponse, error) {
	taken := s.ac.EmailTaken(ctx, req.Email)

	return &ag.EmailTakenResponse{Taken: taken}, nil
}
