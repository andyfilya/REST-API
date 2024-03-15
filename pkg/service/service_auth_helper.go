package service

import (
	"github.com/sirupsen/logrus"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

const (
	entropyDefault = 40
)

func hashPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error while generate hash from password : [%v]", err)
		return nil
	}

	return hash
}

func isValid(password string) bool {
	err := passwordvalidator.Validate(password, entropyDefault)
	if err != nil {
		return false
	}
	return true
}
