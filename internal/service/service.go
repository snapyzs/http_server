package service

import (
	"http_server/internal/repository"
	"http_server/model"
)

type Auth interface {
	CreateUser(user model.Users) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(tokenAccess string) (int, error)
}

type Lists interface {
	GetLists() ([]model.Lists, error)
	CreateList(title, description string) (int, error)
	GetListById(id int) (model.Lists, error)
	UpdateListById(id int, title, description string) (int, error)
	DeleteListById(id int) error
}

type Service struct {
	Auth
	Lists
}

func NewService(r *repository.Repos) *Service {
	return &Service{
		Auth:  NewAuthService(r.Auth),
		Lists: NewListsService(r.Lists),
	}
}
