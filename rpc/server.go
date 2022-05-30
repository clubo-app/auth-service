package rpc

import (
	"log"
	"net"
	"strings"

	"github.com/clubo-app/auth-service/service"
	ag "github.com/clubo-app/protobuf/auth"
	"google.golang.org/grpc"
)

type authServer struct {
	token service.TokenManager
	pw    service.PasswordManager
	goog  service.GoogleManager
	ac    service.AccountService

	ag.UnimplementedAuthServiceServer
}

func NewAuthServer(token service.TokenManager, pw service.PasswordManager, goog service.GoogleManager, ac service.AccountService) ag.AuthServiceServer {
	return &authServer{
		token: token,
		pw:    pw,
		goog:  goog,
		ac:    ac,
	}
}

func Start(s ag.AuthServiceServer, port string) {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(port)
	conn, err := net.Listen("tcp", sb.String())
	if err != nil {
		log.Fatalln(err)
	}

	grpc := grpc.NewServer()

	ag.RegisterAuthServiceServer(grpc, s)

	log.Println("Starting gRPC Server at: ", sb.String())
	if err := grpc.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
