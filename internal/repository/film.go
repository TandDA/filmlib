package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/TandDA/filmlib/internal/model"
	"github.com/sirupsen/logrus"
)

type FilmRepository struct {
	db *sql.DB
}

func NewFilmRepository(db *sql.DB) *FilmRepository {
	return &FilmRepository{db: db}
}

func (r *FilmRepository) Save(film model.FilmCreate) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	query := "INSERT INTO film(name, description, release_date, rating) VALUES($1,$2,$3,$4) RETURNING id;"
	var id int
	row := tx.QueryRow(query, film.Name, film.Description, film.ReleaseDate.GetString(), film.Rating)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, actorId := range film.Actors { // TODO insert many at one request method
		_, err := tx.Exec("INSERT INTO actor_film(actor_id, film_id) VALUES ($1, $2);", actorId, id)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	return id, tx.Commit()
}
func (r *FilmRepository) Update(film model.Film) error {
	query := "UPDATE film SET name=$1, description=$2, release_date=$3, rating=$4 WHERE id=$5;"
	_, err := r.db.Exec(query, film.Name, film.Description, film.ReleaseDate.GetString(), film.Rating, film.Id)
	return err
}
func (r *FilmRepository) Delete(filmId int) error {
	query := "DELETE FROM film WHERE id=$1;"
	_, err := r.db.Exec(query, filmId)
	return err
}
func (r *FilmRepository) GetByPartialName(filmName, actorName string) ([]model.Film, error) {
	query := `
	SELECT DISTINCT film.id, film.name, description, release_date, rating 
	FROM film
	JOIN actor_film ON film.id = film_id
	JOIN actor ON actor.id = actor_id
	WHERE actor.name LIKE '%'|| $1 ||'%'
	AND film.name LIKE '%'|| $2 ||'%'
	`
	rows, err := r.db.Query(query, actorName, filmName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	films := []model.Film{}
	for rows.Next() {
		film := model.Film{}
		err := rows.Scan(&film.Id, &film.Name, &film.Description, &film.ReleaseDate, &film.Rating)
		if err != nil {
			logrus.Error("Cannot read film from rows")
			continue
		}
		films = append(films, film)
	}
	return films, nil
}
func (r *FilmRepository) GetWithSort(column, direction string) ([]model.Film, error) {
	var query string
	switch strings.ToLower(direction) {
	case "asc":
		query = "SELECT * FROM film ORDER BY %s ASC;"
	default:
		query = "SELECT * FROM film ORDER BY %s DESC;"
	}
	if !validParam(column) {
		return nil, errors.New("SQL injection alert")
	}
	query = fmt.Sprintf(query, column)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	films := []model.Film{}
	for rows.Next() {
		film := model.Film{}
		err := rows.Scan(&film.Id, &film.Name, &film.Description, &film.ReleaseDate, &film.Rating)
		if err != nil {
			logrus.Error("Cannot read film from rows")
			continue
		}
		films = append(films, film)
	}
	return films, nil
}

func (r *FilmRepository) GetByActorName(actorName string) ([]model.Film, error) {
	query := `
	SELECT DISTINCT film.id, film.name, description, release_date, rating 
	FROM film
	JOIN actor_film ON film.id = film_id
	JOIN actor ON actor.id = actor_id
	WHERE actor.name = $1
	`

	rows, err := r.db.Query(query, actorName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	films := []model.Film{}
	for rows.Next() {
		film := model.Film{}
		err := rows.Scan(&film.Id, &film.Name, &film.Description, &film.ReleaseDate, &film.Rating)
		if err != nil {
			logrus.Error("Cannot read film from db")
			continue
		}
		films = append(films, film)
	}
	return films, nil
}
