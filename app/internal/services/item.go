package services

import (
	"testAssignment/internal/domain/item"
	"testAssignment/internal/repositories"
)

type ItemService struct {
	repo repositories.Item
}

func (s ItemService) Create(dto item.CreateItemDTO) (int, error) {
	return s.repo.Create(dto)
}

func (s ItemService) GetAll() ([]item.GetItemsDto, error) {
	return s.repo.GetAll()
}

func (s ItemService) GetById(id int) (item.GetItemsDto, error) {
	return s.repo.GetById(id)
}

func (s ItemService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s ItemService) Update(id int, dto item.UpdateItemDTO) error {
	return s.repo.Update(id, dto)
}

func (s ItemService) AddCategoryForItem(itemId, categoryId int) error {
	return s.repo.AddCategoryForItem(itemId, categoryId)
}

func NewItemService(repo repositories.Item) *ItemService {
	return &ItemService{repo: repo}
}
