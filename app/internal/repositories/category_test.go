package repositories

import (
	"testAssignment/internal/config"
	"testAssignment/internal/domain/category"
	"testAssignment/pkg/client/postgresql"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	db, _ := postgresql.NewPostgresDB(config.GetConfig())
	repo := CategoryRepo{db: db}

	dto := category.CreateCategoryDTO{
		Name: "Test Category",
	}

	id, err := repo.Create(dto)

	if err != nil {
		t.Errorf("error creating category: %v", err)
	}

	if id == 0 {
		t.Errorf("expected non-zero category ID")
	}
}

func TestDeleteCategory(t *testing.T) {
	db, _ := postgresql.NewPostgresDB(config.GetConfig())
	repo := CategoryRepo{db: db}

	id := 1

	err := repo.Delete(id)

	if err != nil {
		t.Errorf("error deleting category: %v", err)
	}
}
