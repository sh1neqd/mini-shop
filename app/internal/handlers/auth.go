package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"testAssignment/internal/domain/user"
)

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	u, err := h.services.Authorization.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var input user.CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid json body"))
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		resp := fmt.Sprintf("failed to create user, error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(resp))
		return
	}

	w.Write([]byte(string(rune(id))))
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var input user.AuthorizeUserDto
	var err error
	err = json.NewDecoder(r.Body).Decode(&input)
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		r := fmt.Sprintf(" error authorize, err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(r))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func (h *Handler) test(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if h.userIdentity(w, r, params) {
		w.Write([]byte("test"))
	}
}
