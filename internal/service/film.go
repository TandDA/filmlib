package service

import (
	"github.com/TandDA/filmlib/internal/model"
	"github.com/TandDA/filmlib/internal/repository"
)

type FilmService struct {
	repo repository.Film
}

func NewFilmService(repo repository.Film) *FilmService {
	return &FilmService{repo: repo}
}

func (s *FilmService) Save(film model.FilmCreate) (int, error) {
	return s.repo.Save(film)
}
func (s *FilmService) Update(film model.Film) error {
	return s.repo.Update(film)
}
func (s *FilmService) Delete(filmId int) error {
	return s.repo.Delete(filmId)
}
func (s *FilmService) GetByName(filmName, actorName string) ([]model.Film, error) {
	return s.repo.GetByName(filmName, actorName)
}
func (s *FilmService) GetWithSort(column, direction string) ([]model.Film, error) {
	return s.repo.GetWithSort(column, direction)
}