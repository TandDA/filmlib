package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/TandDA/filmlib/internal/handler"
	"github.com/TandDA/filmlib/internal/repository"
	"github.com/TandDA/filmlib/internal/service"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Print(err)
		return
	}
	doMigration(db)

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	http.ListenAndServe(":8080", handler.InitRoutes())
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
