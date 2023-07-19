package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/repository"
)

const salt = "m2lm#bgib&mtretibertba6msd,varb!gt"

type AuthService struct {
	repo repository.UserAuthorization
}

func NewAuthService(repo repository.UserAuthorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(u todo.User) (int, error) {
	u.Password = s.generatePasswordHash(u.Password)
	return s.repo.CreateUser(u)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
