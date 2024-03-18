package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/TandDA/filmlib/internal/handler"
	"github.com/TandDA/filmlib/internal/repository"
	"github.com/TandDA/filmlib/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title Filmlib API
// @version 1.0
// @description API Server for films and actors

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	db, err := sql.Open("postgres", "postgres://postgres:123@db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Print(err)
		return
	}
	doMigration(db)
	validate := validator.New()
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service, validate)
	http.ListenAndServe(":8080", handler.InitRoutes()) // TODO graceful shutdown
}

func doMigration(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Print(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Print(err)
		return
	}
	m.Up()
}
