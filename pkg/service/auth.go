package service

import (
	"context"
	"github.com/bibishkin/bi-notes-rest-api/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(ctx context.Context, username, password string) (int, error) {
	passwordHash, err := s.hashPassword(password)
	if err != nil {
		return 0, err
	}
	return s.repo.CreateUser(ctx, username, passwordHash)
}

func (s *AuthService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *AuthService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
