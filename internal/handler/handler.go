package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TandDA/filmlib/internal/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service *service.Service
	validate *validator.Validate
}

func NewHandler(service *service.Service, validate *validator.Validate) *Handler {
	return &Handler{
		service: service,
		validate: validate,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	stdMux := http.NewServeMux()
	stdMux.HandleFunc("/actor/all", h.getAllActors)
	stdMux.HandleFunc("/actor/save", h.saveActor)
	stdMux.HandleFunc("/actor/update", h.updateActor)
	stdMux.HandleFunc("/actor/delete", h.deleteActor)

	stdMux.HandleFunc("/film/all", h.getAllFilmsWithSort)
	stdMux.HandleFunc("/film/name", h.getFilmByName)
	stdMux.HandleFunc("/film/update", h.updateFilm)
	stdMux.HandleFunc("/film/delete", h.deleteFilm)
	stdMux.HandleFunc("/film/save", h.saveFilm)
	return stdMux
}

func returnErr(w http.ResponseWriter, statusCode int, requestErr error) {
	js, err := json.Marshal(map[string]string{"err": requestErr.Error()})
	if err != nil {
		//TODO log
	}
	http.Error(w, string(js), statusCode)
}

func returnJSON(w http.ResponseWriter, v any, statusCode int) {
	js, err := json.Marshal(v)
	if err != nil {
		//TODO log
	}
	w.WriteHeader(statusCode)
	w.Write(js)
}
