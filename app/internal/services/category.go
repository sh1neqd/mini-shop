package services

import (
	"testAssignment/internal/domain/category"
	"testAssignment/internal/repositories"
)

type CategoryService struct {
	repo repositories.Category
}

func (s CategoryService) Create(dto category.CreateCategoryDTO) (int, error) {
	return s.repo.Create(dto)
}

func (s CategoryService) GetAll() ([]category.Category, error) {
	return s.repo.GetAll()
}

func (s CategoryService) GetById(id int) (category.Category, error) {
	return s.repo.GetById(id)
}

func (s CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s CategoryService) Update(id int, dto category.CreateCategoryDTO) error {
	return s.repo.Update(id, dto)
}

func NewCategoryService(repo repositories.Category) *CategoryService {
	return &CategoryService{repo: repo}
}
