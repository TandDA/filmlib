package model

import "time"

type Actor struct {
	Id        int
	Name      string
	Male      bool
	BirthDate time.Time
	Films     []Film
}

type Film struct {
	Id          int
	Name        string
	Description string
	ReleaseDate time.Time
	Rating      int
}

type ActorUpdate struct {
	Id            int
	Name          string
	Male          bool
	BirthDate     time.Time
	AddFilmIds    []int
	DeleteFilmIds []int
}
