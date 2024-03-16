package repository

import (
	"database/sql"

	"github.com/TandDA/filmlib/internal/model"
)

type Actor interface {
	Save(actor model.Actor) (int, error)
	Update(actor model.ActorUpdate) error
	Delete(actorId int) error
	GetAll() ([]model.Actor, error)
}

type Film interface {
	Save(actor model.Film) (int, error)
	Update(actor model.Film) error
	Delete(filmId int) error
	GetByName(filmName, actorName string) ([]Film, error)
	GetWithSort() ([]Film, error)
}

type Repository struct {
	Actor
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Actor: NewActorRepository(db),
	}
}
