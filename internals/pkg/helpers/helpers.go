package helpers

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Page    int    `json:"page,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func StatusOK(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusOK, Message{
		Status: "success",
		Data:   data,
	})
}

func StatusCreated(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusCreated, Message{
		Status: "success",
		Data:   data,
	})
}
func StatusAcceptedData(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusAccepted, Message{
		Status: "success",
		Data:   data,
	})
}

func StatusAcceptedMsg(w http.ResponseWriter, msg string) {
	WriteJson(w, http.StatusAccepted, Message{
		Status:  "success",
		Message: msg,
	})
}

func StatusOKAll(w http.ResponseWriter, limit, page int, data any) {
	WriteJson(w, http.StatusOK, Message{
		Status: "success",
		Limit:  limit,
		Page:   page,
		Data:   data,
	})
}

func StatusNotFound(w http.ResponseWriter, err string) {
	WriteJson(w, http.StatusNotFound, Message{
		Status:  "error",
		Message: err,
	})
}

func StatusInternalServerError(w http.ResponseWriter, err string) {
	WriteJson(w, http.StatusInternalServerError, Message{
		Status:  "error",
		Message: err,
	})
}

func StatusBadRequest(w http.ResponseWriter, err string) {
	WriteJson(w, http.StatusBadRequest, Message{
		Status:  "error",
		Message: err,
	})
}
