package rpc

import (
	"github.com/clubo-app/auth-service/service"
	ag "github.com/clubo-app/protobuf/auth"
)

type authServer struct {
	token  service.TokenManager
	pw     service.PasswordManager
	google service.GoogleManager
	ac     service.AccountService

	ag.UnimplementedAuthServiceServer
}

func NewAuthServer() ag.AuthServiceServer {
	return &authServer{}
}
