package service

import (
	"errors"
	"github.com/andyfilya/restapi"
	"github.com/andyfilya/restapi/pkg/repo"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"

	_ "github.com/lib/pq"
)

const (
	timeExpires = time.Hour * 24 * 7 // one week
	signKey     = "fjfuiNDDHAnfklSIFIKDmnNddHmslsoJFFndhI"
)

type AuthService struct {
	repo repo.Authorization
}

func InitAuthService(repo repo.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (auth *AuthService) NewUser(user restapi.User) (int, error) {

	if !isValid(user.Password) {
		logrus.Infof("not valid password for user: [%s]", user.Username)
		return -1, errors.New("not valid password, make it right")
	}
	if user.Role != "user" {
		return -1, errors.New("you are not admin")
	}
	hashedPassword := hashPassword(user.Password)
	user.Password = string(hashedPassword)

	return auth.repo.NewUser(user)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	UserRole string `json:"user_role"`
}

func (auth *AuthService) NewUserToken(username, password string) (string, error) {
	us, err := auth.repo.FindUser(username, password)
	if err != nil {
		logrus.Errorf("error while get new user token : [%v]", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeExpires).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		us.Id,
		us.Role,
	})

	return token.SignedString([]byte(signKey))
}

func (auth *AuthService) ParseUserToken(authToken string) (int, error) {
	token, err := jwt.ParseWithClaims(authToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad auth token")
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("not type of token claim like our service.")
	}

	return claims.UserId, nil
}

func (auth *AuthService) ParseAdminToken(authToken string) (string, error) {
	token, err := jwt.ParseWithClaims(authToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad auth token")
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("not type of token claim like our service")
	}
	return claims.UserRole, nil
}
