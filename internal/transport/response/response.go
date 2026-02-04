package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool      `json:"success"`
	Data    any       `json:"data,omitempty"`
	Error   *APIError `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, status int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
	// TODO: add logger here
}

func Success(w http.ResponseWriter, data any) {
	JSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
	})
}
