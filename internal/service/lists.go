package service

import (
	"http_server/internal/repository"
	"http_server/model"
)

type ItemListService struct {
	repo repository.Lists
}

func NewListsService(r repository.Lists) *ItemListService {
	return &ItemListService{repo: r}
}

func (i *ItemListService) GetLists() ([]model.Lists, error) {
	return i.repo.GetLists()
}

func (i *ItemListService) CreateList(title, description string) (int, error) {
	return i.repo.CreateList(title, description)
}

func (i *ItemListService) GetListById(id int) (model.Lists, error) {
	return i.repo.GetListById(id)
}

func (i *ItemListService) UpdateListById(id int, title, description string) (int, error) {
	return i.repo.UpdateListById(id, title, description)
}

func (i *ItemListService) DeleteListById(id int) error {
	return i.repo.DeleteListById(id)
}
