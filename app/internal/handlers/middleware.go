package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request, params httprouter.Params) bool {
	header := r.Header.Get(authorizationHeader)
	if header == "" {
		newResponse(w, http.StatusUnauthorized, "empty auth header")
		return false
	}
	fmt.Println(header)
	headerParts := strings.Split(header, ".")
	fmt.Println(len(headerParts))
	if len(headerParts) != 3 {
		newResponse(w, http.StatusUnauthorized, "invalid auth header")
		return false
	}

	if len(headerParts[1]) == 0 {
		newResponse(w, http.StatusUnauthorized, "token is empty")
		return false
	}

	userId, err := h.services.Authorization.ParseToken(header)
	if err != nil {
		newResponse(w, http.StatusUnauthorized, err.Error())
		return false
	}
	w.Header().Set(userCtx, fmt.Sprint(userId))
	return true
}
