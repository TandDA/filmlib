package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TandDA/filmlib/internal/model"
)

var sortValues = map[string]struct{}{
	"release_date": {},
	"name":         {},
	"rating":       {},
}
var directionValues = map[string]struct{}{
	"asc":  {},
	"desc": {},
}
// @Summary Save film
// @Security ApiKeyAuth
// @Description Save a film to the database
// @Tags Film
// @Accept json
// @Produce json
// @Param film body model.FilmCreate true "Film data to save"
// @Success 201 {object} idStruct "Film saved successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /film/save [post]
func (h *Handler) saveFilm(w http.ResponseWriter, r *http.Request) {
	var film model.FilmCreate
	err := json.NewDecoder(r.Body).Decode(&film)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	err = h.validate.Struct(film)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.Film.Save(film)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	returnJSON(w, idStruct{id}, http.StatusCreated)
}
// @Summary Update a film
// @Security ApiKeyAuth
// @Description Update a film in the database
// @Tags Film
// @Accept json
// @Produce json
// @Param film body model.Film true "Film object that needs to be updated"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /film/update [put]
func (h *Handler) updateFilm(w http.ResponseWriter, r *http.Request) {
	var updFilm model.Film
	err := json.NewDecoder(r.Body).Decode(&updFilm)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	err = h.validate.Struct(updFilm)
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
// @Summary Delete a film by ID
// @Security ApiKeyAuth
// @Description Delete a film from the database by its ID
// @Tags Film
// @Accept json
// @Produce json
// @Param filmId body idStruct true "Film ID to delete"
// @Success 200 {string} string "Film deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /film/delete [delete]
func (h *Handler) deleteFilm(w http.ResponseWriter, r *http.Request) {
	var filmId idStruct
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
// @Summary Get films by partial name and actor
// @Security ApiKeyAuth
// @Tags Film
// @Description Get films by providing a partial film name and actor name
// @Produce json
// @Param actor query string false "Actor name"
// @Param film query string false "Partial film name"
// @Success 200 {array} model.Film "List of films"
// @Failure 500 {object} error "Internal Server Error"
// @Router /film/name [get]
func (h *Handler) getFilmByName(w http.ResponseWriter, r *http.Request) {
	actorName := r.URL.Query().Get("actor")
	filmName := r.URL.Query().Get("film")

	films, err := h.service.Film.GetByPartialName(filmName, actorName)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	returnJSON(w, films, http.StatusOK)
}
// @Summary Get all films with sorting
// @Security ApiKeyAuth
// @Tags Film
// @Description Get all films with the specified sorting parameters
// @Produce json
// @Param sort query string false "Sort films by: [rating, name, release_date]"
// @Param direction query string false "Sort direction: [asc, desc]"
// @Success 200 {array} []model.Film
// @Failure 500 {object} error
// @Router /film/all [get]
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
