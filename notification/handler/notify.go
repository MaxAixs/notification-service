package handler

import (
	"net/http"
	"notification-service/notification"
)

func (h *Handler) IncomingDeadlineNotifications(w http.ResponseWriter, r *http.Request) {
	var inputData []notification.DeadlineUserInfo

	if err := parseJSONBody(w, r, &inputData); err != nil {
		return
	}

	err := h.service.ProcessDeadlineData(inputData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
