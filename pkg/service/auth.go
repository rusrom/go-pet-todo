package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	todo "github.com/rusrom/yt-todo"
	"github.com/rusrom/yt-todo/pkg/repository"
	"time"
)

const (
	salt           = "m2lm#bgib&mtretibertba6msd,varb!gt"
	tokenSignedKey = "jhuerh4u2h624g24f5y2ft4"
	tokenTTL       = 6 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"userId"`
}

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

func (s *AuthService) GenerateJWT(c todo.SignInInput) (string, error) {
	user, err := s.repo.GetUser(c.Username, s.generatePasswordHash(c.Password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})

	return token.SignedString([]byte(tokenSignedKey))
}
