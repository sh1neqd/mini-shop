package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
	"testAssignment/internal/domain/item"
)

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.userIdentity(w, r, params) {
		return
	}
	var input item.CreateItemDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid json body")
		return
	}
	id, err := h.services.Item.Create(input)
	if err != nil {
		newResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to create item, err:%v", err))
		return
	}
	fmt.Println(id)
	newResponse(w, http.StatusCreated, fmt.Sprint(id))
}

func (h *Handler) GetAllItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	items, _ := h.services.Item.GetAll()
	fmt.Println(items)
	jsonResponse, _ := json.Marshal(items)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) GetByItemId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	idStr := segments[len(segments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid id")
		return
	}
	u, err := h.services.Item.GetById(id)
	if err != nil {
		newResponse(w, http.StatusNotFound, "item with this id not found")
		return
	}
	fmt.Println(u)
	response, _ := json.Marshal(u)
	w.WriteHeader(http.StatusFound)
	w.Write(response)
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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
	err = h.services.Item.Delete(id)
	if err != nil {
		newResponse(w, http.StatusNotFound, "item with this id not found")
		return
	}
	newResponse(w, http.StatusOK, "ok")
}

func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.userIdentity(w, r, params) {
		return
	}
	var input item.UpdateItemDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	path := r.URL.Path
	segments := strings.Split(path, "/")
	idStr := segments[len(segments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid id")
		return
	}
	h.services.Item.Update(id, input)
	if err != nil {
		newResponse(w, http.StatusNotFound, fmt.Sprintf("failed to update item, err:%v", err))
		return
	}
	newResponse(w, http.StatusOK, "ok")
}

func (h *Handler) AddCategoryToItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !h.userIdentity(w, r, params) {
		return
	}
	var input item.GetCategoryID
	err := json.NewDecoder(r.Body).Decode(&input)
	path := r.URL.Path
	segments := strings.Split(path, "/")
	idStr := segments[len(segments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newResponse(w, http.StatusBadRequest, "invalid item id")
		return
	}
	err = h.services.Item.AddCategoryForItem(id, input.CategoryId)
	if err != nil {
		newResponse(w, http.StatusNotFound, "item or category with this ids not found")
		return
	}
	newResponse(w, http.StatusOK, "ok")
}
