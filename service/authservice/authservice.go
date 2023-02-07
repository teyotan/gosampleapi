package authservice

import (
	"bytes"
	"crypto/ed25519"
	"gosampleapi/config"
	"gosampleapi/model"
	"gosampleapi/store/mysqlstore"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	s  *mysqlstore.MySQLStore
	c  *config.Config
	pk ed25519.PrivateKey
}

func NewAuthService(s *mysqlstore.MySQLStore, c *config.Config) *AuthService {
	_, pk, err := ed25519.GenerateKey(bytes.NewReader([]byte(c.Secret)))
	if err != nil {
		log.Fatalln("secret unuseable")
		return nil
	}

	return &AuthService{
		s:  s,
		c:  c,
		pk: pk,
	}
}

func (svc *AuthService) GetPublicKey() interface{} {
	return svc.pk.Public()
}

func (svc *AuthService) RegisterUser(username, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return svc.s.CreateUser(username, hashedPassword)
}

// func (svc *AuthService) ValidateToken(token string) (bool, error) {

// }

func (svc *AuthService) GenerateAccessToken(username, password string) (string, error) {
	user, err := svc.s.GetUserWithPasswordByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(svc.pk)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
