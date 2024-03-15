package repository

import (
	"database/sql"

	"github.com/TandDA/filmlib/internal/model"
)

type Actor interface {
	Save(actor model.Actor) (int, error)
}

type Repository struct {
}

func NewRepository(db *sql.DB) {

}
