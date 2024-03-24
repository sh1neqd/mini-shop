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
	// Auth logic
	router.POST("/api/signup", h.signUp)
	router.POST("/api/signin", h.signIn)
	// Category logic
	router.POST("/api/category", h.Create)
	router.GET("/api/category/:id", h.GetById)
	router.PATCH("/api/category/:id", h.Update)
	router.DELETE("/api/category/:id", h.Delete)
	// Item logic
	router.POST("/api/item", h.CreateItem)
	router.GET("/api/item/:id", h.GetByItemId)
	router.PATCH("/api/item/:id", h.UpdateItem)
	router.DELETE("/api/item/:id", h.DeleteItem)
	router.POST("/api/item/:id", h.AddCategoryToItem)
}
