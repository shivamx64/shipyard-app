package handlers

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Version string
}

type Response struct {
	Service string `json:"service,omitempty"`
	Status  string `json:"status,omitempty"`
	Version string `json:"version,omitempty"`
}

func New(version string) *Handler {
	return &Handler{
		Version: version,
	}
}

func (h *Handler) Root(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Service: "shipyard-api",
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status: "healthy",
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *Handler) VersionHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Version: h.Version,
	}

	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}