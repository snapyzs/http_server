package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Port       string
	PortDB     string
	HostDB     string
	NameDB     string
	PasswordDB string
	UsernameDB string
	SSLModeDB  string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}
	cfg := &Config{
		Port:       os.Getenv("PORT"),
		PortDB:     os.Getenv("PORT_DB"),
		HostDB:     os.Getenv("HOST_DB"),
		PasswordDB: os.Getenv("PASS_DB"),
		UsernameDB: os.Getenv("USER_DB"),
		SSLModeDB:  os.Getenv("SSLMODE"),
		NameDB:     os.Getenv("NAME_DB"),
	}
	return cfg
}
