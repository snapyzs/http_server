package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"http_server/config"
)

const (
	userTable = "users"
	listTalbe = "lists"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.HostDB, cfg.PortDB, cfg.UsernameDB, cfg.NameDB, cfg.PasswordDB, cfg.SSLModeDB))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
