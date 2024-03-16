package service

import (
	"github.com/TandDA/filmlib/internal/model"
	"github.com/TandDA/filmlib/internal/repository"
)

type ActorService struct {
	repo repository.Actor
}

func NewActorService(repo repository.Actor) *ActorService {
	return &ActorService{repo: repo}
}

func (s *ActorService) Save(actor model.Actor) (int, error) {
	return s.repo.Save(actor)
}
func (s *ActorService) Update(actor model.ActorUpdate) error {
	return s.repo.Update(actor)
}
func (s *ActorService) Delete(actorId int) error {
	return s.repo.Delete(actorId)
}
func (s *ActorService) GetAll() ([]model.Actor, error) {
	return s.repo.GetAll()
}
