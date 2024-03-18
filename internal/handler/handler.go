package handler

import (
	"encoding/json"
	"net/http"

	_ "github.com/TandDA/filmlib/docs"
	"github.com/TandDA/filmlib/internal/middleware"
	"github.com/TandDA/filmlib/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handler struct {
	service  *service.Service
	validate *validator.Validate
}

func NewHandler(service *service.Service, validate *validator.Validate) *Handler {
	return &Handler{
		service:  service,
		validate: validate,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	admMux := http.NewServeMux()
	admMux.HandleFunc("/actor/save", h.saveActor)
	admMux.HandleFunc("/actor/update", h.updateActor)
	admMux.HandleFunc("/actor/delete", h.deleteActor)
	admMux.HandleFunc("/film/update", h.updateFilm)
	admMux.HandleFunc("/film/delete", h.deleteFilm)
	admMux.HandleFunc("/film/save", h.saveFilm)

	userMux := http.NewServeMux()
	userMux.HandleFunc("/actor/all", h.getAllActors)

	userMux.HandleFunc("/film/all", h.getAllFilmsWithSort)
	userMux.HandleFunc("/film/name", h.getFilmByName)

	userMux.Handle("/", middleware.AdminAuth(admMux))
	userMux.HandleFunc("/user/auth", h.AuthUser)

	openMux := http.NewServeMux()
	openMux.Handle("/", middleware.UserAuth(userMux))
	openMux.HandleFunc("/user/auth", h.AuthUser)
	openMux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	siteHandler := middleware.Logger(openMux)
	return siteHandler
}

func returnErr(w http.ResponseWriter, statusCode int, requestErr error) {
	js, err := json.Marshal(map[string]string{"err": requestErr.Error()})
	if err != nil {
		logrus.Error("Cannot convert error to json")
	}
	http.Error(w, string(js), statusCode)
}

func returnJSON(w http.ResponseWriter, v any, statusCode int) {
	js, err := json.Marshal(v)
	if err != nil {
		logrus.Error("Cannot convert object to json")
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}
