package repositories

import (
	"github.com/jmoiron/sqlx"
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
	Create() (int, error)
	GetAll() error
	GetById()
	Delete(userId, listId int) error
}

type Category interface {
	Create()
	GetAll(userId, listId int)
	GetById(userId, itemId int)
	Delete(userId, itemId int) error
	Update(userId, itemId int)
}

type Repository struct {
	Authorization
	//Item
	//Category
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
		//Item:          NewItemPostgres(db),
		//Category:      NewCategoryPostgres(db),
	}
}
