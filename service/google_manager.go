package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/clubo-app/auth-service/datastruct"
	"github.com/golang-jwt/jwt"
)

type GoogleManager interface {
	ValidateGoogleJWT(token string) (datastruct.GoogleClaims, error)
	getGooglePublicKey(keyID string) (string, error)
}

type googleManager struct {
	clientId string
}

func NewGoogleManager(clientId string) googleManager {
	return googleManager{clientId: clientId}
}

func (g googleManager) ValidateGoogleJWT(tokenString string) (datastruct.GoogleClaims, error) {
	claimsStruct := datastruct.GoogleClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (any, error) {
			pem, err := g.getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return datastruct.GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*datastruct.GoogleClaims)
	if !ok {
		return datastruct.GoogleClaims{}, errors.New("invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "accounts.google.com" {
		return datastruct.GoogleClaims{}, errors.New("iss is invalid")
	}

	if claims.Audience != g.clientId {
		return datastruct.GoogleClaims{}, errors.New("aud is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return datastruct.GoogleClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}

func (g googleManager) getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}
