package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"http_server/model"
)

type ItemListRepos struct {
	db *sqlx.DB
}

func NewItemList(db *sqlx.DB) *ItemListRepos {
	return &ItemListRepos{db: db}
}

func (r *ItemListRepos) GetLists() ([]model.Lists, error) {
	var lists []model.Lists
	query := fmt.Sprintf("SELECT id,title,description FROM %s", listTalbe)
	err := r.db.Select(&lists, query)
	return lists, err
}

func (r *ItemListRepos) CreateList(title, description string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title,description) values ($1,$2) RETURNING id", listTalbe)
	row := r.db.QueryRow(query, title, description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *ItemListRepos) GetListById(id int) (model.Lists, error) {
	var lists model.Lists
	query := fmt.Sprintf("SELECT id,title,description FROM %s WHERE id=$1", listTalbe)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&lists.Id, &lists.Title, &lists.Description); err != nil {
		return lists, err
	}
	return lists, nil
}

func (r *ItemListRepos) UpdateListById(id int, title, description string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := fmt.Sprintf("UPDATE %s SET (title,description) = ($1,$2) WHERE id = $3 RETURNING id", listTalbe)
	row := r.db.QueryRow(query, title, description, id)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *ItemListRepos) DeleteListById(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", listTalbe)
	if _, err := r.db.Query(query, id); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
