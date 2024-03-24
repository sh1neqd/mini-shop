package repositories

import (
	"github.com/jmoiron/sqlx"
	"testAssignment/internal/domain/category"
	"testAssignment/internal/domain/item"
	"testAssignment/internal/domain/user"
)

type Authorization interface {
	CreateUser(dto user.CreateUserDTO) (int, error)
	GetById(id int) (user.User, error)
	GetUser(username, password string) (user.User, error)
	PasswordsPass(username, password string) bool
	UserExist(username string) bool
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

type Repository struct {
	Authorization
	Item
	Category
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		Item:          NewItemRepo(db),
		Category:      NewCategoryRepo(db),
	}
}
