package repository

import (
	"database/sql"

	"github.com/TandDA/filmlib/internal/model"
	"github.com/sirupsen/logrus"
)

type ActorRepository struct {
	db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{db: db}
}

func (r *ActorRepository) Save(actor model.Actor) (int, error) {
	query := "INSERT INTO actor(name,male,birth_date) VALUES($1,$2,$3) RETURNING id;"
	var id int
	row := r.db.QueryRow(query, actor.Name, actor.Male, actor.BirthDate.GetString())
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ActorRepository) Update(actor model.ActorUpdate) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	actorQuery := "UPDATE actor SET name=$1, male=$2, birth_date=$3 WHERE id=$4;"
	_, err = tx.Exec(actorQuery, actor.Name, actor.Male, actor.BirthDate.GetString(), actor.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, addId := range actor.AddFilmIds {
		_, err := tx.Exec("INSERT INTO actor_film(actor_id, film_id) VALUES ($1, $2);", actor.Id, addId) // TODO insert many method
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for _, delId := range actor.AddFilmIds {
		_, err := tx.Exec("DELETE FROM actor_film(actor_id, film_id) WHERE actor_id=$1 AND film_id=$2;", actor.Id, delId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *ActorRepository) Delete(actorId int) error {
	query := "DELETE FROM actor WHERE id=$1;"
	_, err := r.db.Exec(query, actorId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ActorRepository) GetAll() ([]model.Actor, error) {
	actorQuery := "SELECT * FROM actor;"
	actorRows, err := r.db.Query(actorQuery)
	if err != nil {
		return nil, err
	}
	defer actorRows.Close()

	actors := []model.Actor{}
	for actorRows.Next() {
		actor := model.Actor{}
		err := actorRows.Scan(&actor.Id, &actor.Name, &actor.Male, &actor.BirthDate)
		if err != nil {
			logrus.Error("Cannot read actor details")
			continue
		}
		actors = append(actors, actor)
	}
	return actors, nil
}
