package main

import (
	"log"

	"github.com/clubo-app/auth-service/config"
	"github.com/clubo-app/auth-service/repository"
	"github.com/clubo-app/auth-service/rpc"
	"github.com/clubo-app/auth-service/service"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	r, err := repository.NewAccountRepository(c.DB_USER, c.DB_PW, c.DB_NAME, c.DB_HOST, c.DB_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	t := service.NewTokenManager(c.TOKEN_SECRET)
	goog := service.NewGoogleManager(c.GOOGLE_CLIENTID)
	pw := service.NewPasswordManager()

	as := service.NewAccountService(r)

	s := rpc.NewAuthServer(t, pw, goog, as)

	rpc.Start(s, c.PORT)
}
