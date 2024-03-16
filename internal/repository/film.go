package repository

import (
	"database/sql"

	"github.com/TandDA/filmlib/internal/model"
)

type FilmRepository struct {
	db *sql.DB
}

func (r *FilmRepository) Save(film model.Film) (int, error) {
	query := "INSERT INTO film(name, description, release_date, rating) VALUES($1,$2,$3,$5) RETURNING id;"
	var id int
	row := r.db.QueryRow(query, film.Name, film.Description, film.ReleaseDate, film.Rating)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *FilmRepository) Update(film model.Film) error {
	return nil
}
func (r *FilmRepository) Delete(filmId int) error {
	return nil
}
func (r *FilmRepository) GetByName(filmName, actorName string) ([]Film, error) {
	return nil, nil
}
func (r *FilmRepository) GetWithSort() ([]Film, error) {
	return nil, nil
}
