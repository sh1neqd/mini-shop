package repositories

import (
	"testAssignment/internal/config"
	"testAssignment/internal/domain/item"
	"testAssignment/pkg/client/postgresql"
	"testing"
)

func TestCreateItem(t *testing.T) {
	db, _ := postgresql.NewPostgresDB(config.GetConfig())
	repo := ItemRepo{db: db}

	dto := item.CreateItemDTO{
		Name:  "Test Item",
		Price: 10,
	}

	id, err := repo.Create(dto)

	if err != nil {
		t.Errorf("error creating item: %v", err)
	}

	if id == 0 {
		t.Errorf("expected non-zero item ID")
	}
}

func TestGetItemById(t *testing.T) {
	db, _ := postgresql.NewPostgresDB(config.GetConfig())
	repo := ItemRepo{db: db}

	id := 1

	item, err := repo.GetById(id)

	if err != nil {
		t.Errorf("Error getting item by ID: %v", err)
	}

	if item.Name == "" {
		t.Errorf("Expected non-empty item name")
	}

	if item.Price == 0 {
		t.Errorf("Expected non-zero item price")
	}
}
