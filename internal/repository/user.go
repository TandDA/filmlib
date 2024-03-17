package repository

import (
	"database/sql"

	"github.com/TandDA/filmlib/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByEmail(email string) (model.User, error) {
	query := "SELECT * FROM users WHERE email=$1;"
	var user model.User
	row := r.db.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.RoleId, &user.Email, &user.Password)
	return user, err
}
