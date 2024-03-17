package services

import (
	"testAssignment/internal/domain/user"
	"testAssignment/internal/repositories"
)

type Authorization interface {
	CreateUser(dto user.CreateUserDTO) (int, error)
	GetById(id int) (user.User, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
