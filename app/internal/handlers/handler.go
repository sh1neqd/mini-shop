package handlers

import (
	"github.com/julienschmidt/httprouter"
	"testAssignment/internal/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(router *httprouter.Router) {
	router.POST("/api/", h.signUp)
	router.POST("/api/login", h.signIn)
	router.GET("/api/", h.test)
}
