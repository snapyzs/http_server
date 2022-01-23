package repository

import (
	"github.com/jmoiron/sqlx"
	"http_server/model"
)

type Auth interface {
	CreateUser(user model.Users) (int, error)
	GetUser(email, password string) (model.Users, error)
}

type Lists interface {
	GetLists() ([]model.Lists, error)
	CreateList(title, description string) (int, error)
	GetListById(id int) (model.Lists, error)
	UpdateListById(id int, title, description string) (int, error)
	DeleteListById(id int) error
}

type Repos struct {
	Auth
	Lists
}

func NewRepository(db *sqlx.DB) *Repos {
	return &Repos{
		Auth:  NewAuthPostgres(db),
		Lists: NewItemList(db),
	}
}
