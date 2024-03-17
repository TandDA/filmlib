package service

import (
	"github.com/TandDA/filmlib/internal/model"
	"github.com/TandDA/filmlib/internal/repository"
)

type ActorService struct {
	actorRepo repository.Actor
	filmRepo  repository.Film
}

func NewActorService(actorRepo repository.Actor, filmRepo repository.Film) *ActorService {
	return &ActorService{actorRepo: actorRepo, filmRepo: filmRepo}
}

func (s *ActorService) Save(actor model.Actor) (int, error) {
	return s.actorRepo.Save(actor)
}
func (s *ActorService) Update(actor model.ActorUpdate) error {
	return s.actorRepo.Update(actor)
}
func (s *ActorService) Delete(actorId int) error {
	return s.actorRepo.Delete(actorId)
}
func (s *ActorService) GetAll() ([]model.Actor, error) {
	actors, err := s.actorRepo.GetAll()
	if err != nil {
		return nil, err
	}
	for i := range actors {
		films, err := s.filmRepo.GetByActorName(actors[i].Name)
		if err != nil {
			return nil, err
		}
		actors[i].Films = append(actors[i].Films, films...)
	}
	return actors, nil
}
