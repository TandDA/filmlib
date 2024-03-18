package handler

import (
	"encoding/json"
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
// @Failure 400 {object} errorResponse "Failed to decode request body. Invalid JSON"
// @Failure 500 {object} errorResponse "Failed to generate JWT"
// @Router /user/auth [post]
func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {
	requestUser := signInNpit{}
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		returnErr(w, http.StatusBadRequest, "Failed to decode request body. Invalid JSON")
		return
	}
	dbUser, err := h.service.GetByEmail(requestUser.Email)
	if err != nil || dbUser.Password != requestUser.Password {
		returnErr(w, http.StatusBadRequest, "Failed to decode request body. Invalid JSON")
		return
	}
	jwtStr, err := util.GenerateJWT(dbUser)
	if err != nil {
		returnErr(w, http.StatusInternalServerError, "Failed to generate JWT")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	returnJSON(w, jwtResponse{jwtStr}, http.StatusOK)
}
