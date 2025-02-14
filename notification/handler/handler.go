package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"notification-service/notification/service"
)

type Handler struct {
	service *service.NotificationService
}

func NewHandler(service *service.NotificationService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) MapRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/notify", h.IncomingDeadlineNotifications).Methods("POST")

	return r
}
