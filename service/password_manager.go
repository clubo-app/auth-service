package service

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordManager interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type passwordManager struct{}

func NewPasswordManager() PasswordManager {
	return &passwordManager{}
}

func (p passwordManager) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (p passwordManager) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
