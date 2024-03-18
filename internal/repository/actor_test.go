package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/TandDA/filmlib/internal/model"
)

func ActorSaveMethod(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewActorRepository(db)

	actor := model.Actor{
		Name:      "John Doe",
		Male:      true,
		BirthDate: model.Date(time.Now()),
	}
	mock.ExpectQuery("INSERT INTO actor").
		WithArgs(actor.Name, actor.Male, actor.BirthDate.GetString()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	_, err = repo.Save(actor)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	//query error
	mock.ExpectQuery("INSERT INTO actor").
		WithArgs(actor.Name, actor.Male, actor.BirthDate.GetString()).
		WillReturnError(fmt.Errorf("bad request"))

	_, err = repo.Save(actor)
	if err == nil {
		t.Errorf("Expecred error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
