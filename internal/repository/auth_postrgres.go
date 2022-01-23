package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"http_server/model"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.Users) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash) values ($1, $2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *AuthPostgres) GetUser(email, password string) (model.Users, error) {
	var user model.Users
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}
