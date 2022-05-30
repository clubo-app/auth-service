package dto

import (
	"github.com/clubo-app/auth-service/repository"
)

type Account struct {
	ID            string
	Email         string
	EmailVerified bool
	EmailCode     string
	PasswordHash  string
	Provider      repository.Provider
	Role          repository.Role
	Type          repository.Type
}
