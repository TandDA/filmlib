package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TandDA/filmlib/internal/model"
	"github.com/TandDA/filmlib/internal/util"
)

func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {
	requestUser := model.User{}
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		returnErr(w, http.StatusBadRequest, err)
		return
	}
	dbUser, err := h.service.GetByEmail(requestUser.Email)
	if err != nil || dbUser.Password != requestUser.Password {
		returnErr(w, http.StatusBadRequest, errors.New("failed authentication"))
		return
	}
	jwtStr, err := util.GenerateJWT(dbUser)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, err)
		return
	}
	returnJSON(w, map[string]string{"jwt": jwtStr}, http.StatusOK)
}
