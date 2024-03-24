package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
	"testAssignment/internal/domain/category"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.userIdentity(w, r, params) {
		return
	}
	var input category.CreateCategoryDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid json body")
		return
	}
	id, err := h.services.Category.Create(input)
	if err != nil {
		newResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to create category, err:%v", err))
		return
	}
	fmt.Println(id)
	newResponse(w, http.StatusCreated, fmt.Sprint(id))
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categories, _ := h.services.Category.GetAll()
	fmt.Println(categories)
	jsonResponse, _ := json.Marshal(categories)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	idStr := segments[len(segments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid id")
		return
	}
	u, err := h.services.Category.GetById(id)
	if err != nil {
		newResponse(w, http.StatusNotFound, "category with this id not found")
		return
	}
	fmt.Println(u)
	response, _ := json.Marshal(u)
	w.WriteHeader(http.StatusFound)
	w.Write(response)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.userIdentity(w, r, params) {
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	idStr := segments[len(segments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.services.Category.Delete(id)
	if err != nil {
		newResponse(w, http.StatusNotFound, "category with this id not found")
		return
	}
	newResponse(w, http.StatusOK, "ok")
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.userIdentity(w, r, params) {
		return
	}
	var input category.CreateCategoryDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	path := r.URL.Path
	segments := strings.Split(path, "/")
	idStr := segments[len(segments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid id")
		return
	}
	h.services.Category.Update(id, input)
	if err != nil {
		newResponse(w, http.StatusNotFound, fmt.Sprintf("failed to update category, err:%v", err))
		return
	}
	newResponse(w, http.StatusOK, "ok")
}
