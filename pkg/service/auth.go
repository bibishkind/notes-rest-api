package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

const jwtTTL = 12 * time.Hour

type JWTClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func (s *Service) CreateUser(username, password string) (int, error) {
	hashedPassword := getHashedPassword(password)

	return s.repository.CreateUser(username, hashedPassword)
}

func (s *Service) GenerateJWT(username, password string) (string, error) {

	user, err := s.repository.GetUser(username, getHashedPassword(password))
	if err != nil {
		return "", err
	}

	jwt, err := generateJWT(user.Id)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (s *Service) ParseJWT(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, errors.New("invalid jwt")
	}
}

func getHashedPassword(password string) string {
	hash := sha256.New()
	bytes := hash.Sum([]byte{})

	hashedPassword := base64.StdEncoding.EncodeToString(bytes)
	return hashedPassword
}

func generateJWT(userId int) (string, error) {
	JWTKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
