package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type AccountRefreshResponse struct {
	AccountID   string    `json:"account_id"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	LastUpdated time.Time `json:"last_updated"`
	NewBalance  *float64  `json:"new_balance,omitempty"`
}

// Handler serves POST /api/accounts/refresh?id=acc_1
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, APIResponse{Success: false, Message: "Account ID is required"})
		return
	}

	newBal := 2750.25
	resp := AccountRefreshResponse{
		AccountID:   id,
		Success:     true,
		Message:     "Account refreshed successfully",
		LastUpdated: time.Now(),
		NewBalance:  &newBal,
	}
	writeJSON(w, http.StatusOK, APIResponse{Success: true, Message: "Account refreshed successfully", Data: resp})
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
