package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testAssignment/internal/domain/user"
)

//func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	path := r.URL.Path
//	segments := strings.Split(path, "/")
//	idStr := segments[len(segments)-1]
//	id, err := strconv.Atoi(idStr)
//	fmt.Println(id)
//	u, err := h.services.Authorization.GetById(id)
//	if err != nil {
//		newResponse(w, http.StatusNotFound, "user with this id not found")
//		return
//	}
//	fmt.Println(u)
//	w.WriteHeader(http.StatusOK)
//	jsonResponse, err := json.Marshal(u)
//	if err != nil {
//		logrus.Errorf("failed to marshal json, err:%v", err)
//		return
//	}
//	w.Write(jsonResponse)
//}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var input user.CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid json body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		resp := fmt.Sprintf("failed to create user, error: %v", err)
		newResponse(w, http.StatusBadRequest, resp)
		return
	}

	newResponse(w, http.StatusOK, fmt.Sprintf("user with id=%d successfully created", id))

}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var input user.AuthorizeUserDto
	var err error
	err = json.NewDecoder(r.Body).Decode(&input)
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		r := fmt.Sprintf("failed authorize, err: %v", err)
		newResponse(w, http.StatusBadRequest, r)
		return
	}

	newResponse(w, http.StatusOK, token)
}
