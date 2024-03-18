package handler

import (
	"encoding/json"
	"net/http"

	"github.com/TandDA/filmlib/internal/model"
)

type saveActorDTO struct {
	Name      string
	Male      bool
	BirthDate model.Date
}

// @Summary Get All Actors
// @Security ApiKeyAuth
// @Tags Actors
// @Description get all actors
// @ID get-all-actors
// @Accept  json
// @Produce  json
// @Success 200 {object} []model.Actor
// @Failure 500 {object} errorResponse "Failed to get all actors"
// @Router /actor/all [get]
func (h *Handler) getAllActors(w http.ResponseWriter, r *http.Request) {
	actors, err := h.service.Actor.GetAll()
	if err != nil {
		returnErr(w, http.StatusInternalServerError, "Failed to get all actors")
		return
	}
	returnJSON(w, actors, http.StatusOK)
}

// @Summary Save an actor
// @Security ApiKeyAuth
// @Description Save the details of an actor
// @Accept  json
// @Produce json
// @Tags Actors
// @Param actor body saveActorDTO true "Actor object to be saved"
// @Success 201 {object} idStruct "Returns the ID of the saved actor"
// @Failure 400 {object} errorResponse "Failed to decode request body. Invalid JSON"
// @Failure 500 {object} errorResponse "Failed to save actor"
// @Router /actor/save [post]
func (h *Handler) saveActor(w http.ResponseWriter, r *http.Request) {
	var actor saveActorDTO
	err := json.NewDecoder(r.Body).Decode(&actor)
	if err != nil {
		returnErr(w, http.StatusBadRequest, "Failed to decode request body. Invalid JSON")
		return
	}
	id, err := h.service.Actor.Save(model.Actor{
		Name:      actor.Name,
		Male:      actor.Male,
		BirthDate: actor.BirthDate,
	})
	if err != nil {
		returnErr(w, http.StatusInternalServerError, "Failed to save actor")
		return
	}
	returnJSON(w, idStruct{id}, http.StatusCreated)
}

// @Summary Update actor
// @Security ApiKeyAuth
// @Description Update an existing actor
// @Tags Actors
// @Accept json
// @Produce json
// @Param body body model.ActorUpdate true "Actor data to be updated"
// @Success 200
// @Failure 400 {object} errorResponse "Failed to decode request body. Invalid JSON"
// @Failure 500 {object} errorResponse "Failed to update actor"
// @Router /actor/update [put]
func (h *Handler) updateActor(w http.ResponseWriter, r *http.Request) {
	var updActor model.ActorUpdate
	err := json.NewDecoder(r.Body).Decode(&updActor)
	if err != nil {
		returnErr(w, http.StatusBadRequest, "Failed to decode request body. Invalid JSON")
		return
	}
	err = h.service.Actor.Update(updActor)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, "Failed to update actor")
		return
	}
	w.WriteHeader(http.StatusOK)
}

// @Summary Delete an actor
// @Security ApiKeyAuth
// @Description Delete an actor by ID
// @Tags Actors
// @Accept json
// @Produce json
// @Param body body idStruct true "Actor ID to delete"
// @Success 200
// @Failure 400 {object} errorResponse "Failed to decode request body. Invalid JSON"
// @Failure 500 {object} errorResponse "Failed to delete actor"
// @Router /actor/delete [delete]
func (h *Handler) deleteActor(w http.ResponseWriter, r *http.Request) {
	var actorId idStruct
	err := json.NewDecoder(r.Body).Decode(&actorId)
	if err != nil {
		returnErr(w, http.StatusBadRequest, "Failed to decode request body. Invalid JSON")
		return
	}
	err = h.service.Actor.Delete(actorId.Id)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, "Failed to delete actor")
		return
	}
	w.WriteHeader(http.StatusOK)
}
