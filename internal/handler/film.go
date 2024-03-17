package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TandDA/filmlib/internal/model"
)

var sortValues = map[string]struct{} {
	"release_date": {},
	"name":         {},
	"rating":       {},
}
var directionValues = map[string]struct{} {
	"asc": {},
	"desc":         {},
}

func (h *Handler) saveFilm(w http.ResponseWriter, r *http.Request) {
	var film model.FilmCreate
	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.Film.Save(film)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	returnJSON(w, map[string]int{"id": id}, http.StatusCreated)
}
func (h *Handler) updateFilm(w http.ResponseWriter, r *http.Request) {
	var updFilm model.Film
	err := json.NewDecoder(r.Body).Decode(&updFilm)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.Film.Update(updFilm)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) deleteFilm(w http.ResponseWriter, r *http.Request) {
	var filmId struct{ Id int }
	err := json.NewDecoder(r.Body).Decode(&filmId)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.Film.Delete(filmId.Id)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) getFilmByName(w http.ResponseWriter, r *http.Request) {
	actorName := r.URL.Query().Get("actor")
	filmName := r.URL.Query().Get("film")

	films, err := h.service.Film.GetByName(filmName,actorName)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	returnJSON(w, films, http.StatusOK)
}
func (h *Handler) getAllFilmsWithSort(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")
	if _, ok := sortValues[sort]; !ok {
		sort = "rating"
	} 
	direction := r.URL.Query().Get("direction")
	if _, ok := directionValues[direction]; !ok {
		direction = "desc"
	}
	films, err := h.service.Film.GetWithSort(sort, direction)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	returnJSON(w, films, http.StatusOK)
}