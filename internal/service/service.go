package service

import (
	"github.com/TandDA/filmlib/internal/model"
	"github.com/TandDA/filmlib/internal/repository"
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
	GetByPartialName(filmName, actorName string) ([]model.Film, error)
	GetWithSort(column, direction string) ([]model.Film, error)
}

type User interface {
	GetByEmail(email string) (model.User, error)
}

type Service struct {
	Actor
	Film
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Actor: NewActorService(repo.Actor, repo.Film),
		Film:  NewFilmService(repo.Film),
		User:  NewUserService(repo.User),
	}
}
