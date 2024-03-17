package repository

import (
	"database/sql"
	"regexp"

	"github.com/TandDA/filmlib/internal/model"
)

type Actor interface {
	Save(actor model.Actor) (int, error)
	Update(actor model.ActorUpdate) error
	Delete(actorId int) error
	GetAll() ([]model.Actor, error)
}

type Film interface {
	Save(film model.FilmCreate) (int, error)
	Update(film model.Film) error
	Delete(filmId int) error
	GetByName(filmName, actorName string) ([]model.Film, error)
	GetWithSort(column, direction string) ([]model.Film, error)
}

type Repository struct {
	Actor
	Film
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Actor: NewActorRepository(db),
		Film:  NewFilmRepository(db),
	}
}

func validParam(param string) bool {
	valid := regexp.MustCompile("^[A-Za-z0-9_]+$")
	return valid.MatchString(param)
}
