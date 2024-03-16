package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TandDA/filmlib/internal/model"
)

func (h *Handler) getAllActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.service.Actor.GetAll()
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	json, err := json.Marshal(actors)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(json)
}

func (h *Handler) saveActor(w http.ResponseWriter, r *http.Request) {
	var actor model.Actor
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	id, err := h.service.Actor.Save(actor)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	returnJSON(w, map[string]int{"id": id}, http.StatusCreated)
}

func (h *Handler) updateActor(w http.ResponseWriter, r *http.Request) {
	var updActor model.ActorUpdate
	err := json.NewDecoder(r.Body).Decode(&updActor)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.Actor.Update(updActor)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	var actorId struct{ Id int }
	err := json.NewDecoder(r.Body).Decode(&actorId)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.Actor.Delete(actorId.Id)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
