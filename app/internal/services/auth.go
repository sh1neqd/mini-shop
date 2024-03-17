package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"testAssignment/internal/domain/user"
	"testAssignment/internal/repositories"
	"time"
)

type AuthService struct {
	repo repositories.Authorization
}

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s AuthService) CreateUser(dto user.CreateUserDTO) (int, error) {
	dto.Password = generatePasswordHash(dto.Password)
	return s.repo.CreateUser(dto)
}

func (s AuthService) GetById(id int) (user.User, error) {
	return s.repo.GetById(id)
}
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	u, err := s.repo.GetUser(username, generatePasswordHash(password))
	if s.repo.PasswordsPass(username, generatePasswordHash(password)) && s.repo.UserExist(username) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			int(u.ID),
		})
		logrus.Println(token.SignedString([]byte(signingKey)))
		return token.SignedString([]byte(signingKey))
	} else {
		logrus.Errorf("password not passed")
	}
	if err != nil {
		logrus.Errorf("failed to generate token, err:%v", err)
	}
	return "", err

}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func NewAuthService(repo repositories.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
