package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TandDA/filmlib/internal/util"
)

type signInNpit struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type jwtResponse struct {
	JWT string `json:"jwt"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInNpit true "credentials"
// @Success 200 {string} map[string]string "token"
// @Router /user/auth [post]
func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {
	requestUser := signInNpit{}
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
	w.Header().Add("Content-Type", "application/json")
	returnJSON(w, jwtResponse{jwtStr}, http.StatusOK)
}
