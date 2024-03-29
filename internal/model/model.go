package model

import (
	"encoding/json"
	"time"
)

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	RoleId int `json:"role_id"`
}

type Actor struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Male      bool   `json:"male"`
	BirthDate Date   `json:"birth_date" time_format:"2006-01-02"`
	Films     []Film `json:"films"`
}

type Film struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"min=1,max=150"`
	Description string `json:"description" validate:"max=1000"`
	ReleaseDate Date   `json:"release_date"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}

type FilmCreate struct {
	Film
	Actors []int
}

type ActorUpdate struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Male          bool   `json:"male"`
	BirthDate     Date   `json:"birth_date"`
	AddFilmIds    []int  `json:"add_film_ids"`
	DeleteFilmIds []int  `json:"delete_film_ids"`
}

type Date time.Time

func (mt *Date) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = Date(t)
	return nil
}

func (t *Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.GetString() + `"`), nil
}

func (mr *Date) GetString() string {
	return time.Time(*mr).Format("2006-01-02")
}
