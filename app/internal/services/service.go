package services

import (
	"testAssignment/internal/domain/category"
	"testAssignment/internal/domain/item"
	"testAssignment/internal/domain/user"
	"testAssignment/internal/repositories"
)

type Authorization interface {
	CreateUser(dto user.CreateUserDTO) (int, error)
	GetById(id int) (user.User, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Item interface {
	Create(dto item.CreateItemDTO) (int, error)
	GetAll() ([]item.GetItemsDto, error)
	GetById(id int) (item.GetItemsDto, error)
	Delete(id int) error
	Update(id int, dto item.UpdateItemDTO) error
	AddCategoryForItem(itemId, categoryId int) error
}

type Category interface {
	Create(dto category.CreateCategoryDTO) (int, error)
	GetAll() ([]category.Category, error)
	GetById(id int) (category.Category, error)
	Delete(id int) error
	Update(id int, dto category.CreateCategoryDTO) error
}

type Service struct {
	Authorization
	Item
	Category
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Item:          NewItemService(repos.Item),
		Category:      NewCategoryService(repos.Category),
	}
}
