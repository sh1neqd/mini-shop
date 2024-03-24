package repositories

import (
	"testAssignment/internal/config"
	"testAssignment/internal/domain/user"
	"testAssignment/pkg/client/postgresql"
	"testing"
)

func TestCreateUser(t *testing.T) {

	db, err := postgresql.NewPostgresDB(config.GetConfig())
	auth := &Auth{db: db}

	dto := user.CreateUserDTO{
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@example.com",
	}

	id, err := auth.CreateUser(dto)

	if err != nil {
		t.Errorf("failed to create user: %v", err)
	}

	if id == 0 {
		t.Errorf("invalid user id")
	}
}

func TestUserExist(t *testing.T) {
	db, _ := postgresql.NewPostgresDB(config.GetConfig())
	auth := &Auth{db: db}

	username := "test1user"

	exists := auth.UserExist(username)

	if exists {
		t.Errorf("expected user to not exist")
	}
}
