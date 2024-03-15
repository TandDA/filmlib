package repository

import (
	"database/sql"

	"github.com/TandDA/filmlib/internal/model"
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
	row := r.db.QueryRow(query, actor.Name, actor.Male, actor.BirthDate)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
