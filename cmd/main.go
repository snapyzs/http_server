package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"http_server"
	"http_server/config"
	"http_server/internal/hanlder"
	"http_server/internal/repository"
	"http_server/internal/service"
)

//func init() {
//	gin.SetMode(gin.ReleaseMode)
//}

func main() {
	cfg := config.NewConfig()
	server := http_server.Server{}
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := hanlder.NewHandler(services)

	logrus.Printf("Startring server on port %s...", cfg.Port)
	if err := server.Start(":"+cfg.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
