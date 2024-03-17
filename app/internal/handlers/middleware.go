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
		newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
		return false
	}
	fmt.Println(header)
	headerParts := strings.Split(header, ".")
	fmt.Println(len(headerParts))
	if len(headerParts) != 3 {
		newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
		return false
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(w, http.StatusUnauthorized, "token is empty")
		return false
	}

	userId, err := h.services.Authorization.ParseToken(header)
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return false
	}
	w.Header().Set(userCtx, string(rune(userId)))
	return true
}
