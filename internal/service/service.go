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
	Save(actor model.Film) (int, error)
	Update(actor model.Film) error
	Delete(filmId int) error
	GetByName(filmName, actorName string) ([]Film, error)
	GetWithSort() ([]Film, error)
}

type Service struct {
	Actor
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Actor: NewActorService(repo.Actor)}
}
