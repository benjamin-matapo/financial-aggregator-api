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

type Account struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Bank        string    `json:"bank"`
	AccountType string    `json:"account_type"`
	Balance     float64   `json:"balance"`
	Currency    string    `json:"currency"`
	LastUpdated time.Time `json:"last_updated"`
	IsActive    bool      `json:"is_active"`
}

var mockAccounts = []Account{
	{
		ID:          "acc_1",
		Name:        "Main Checking",
		Bank:        "Chase Bank",
		AccountType: "checking",
		Balance:     2500.75,
		Currency:    "USD",
		LastUpdated: time.Now().Add(-24 * time.Hour),
		IsActive:    true,
	},
	{
		ID:          "acc_2",
		Name:        "Savings Account",
		Bank:        "Bank of America",
		AccountType: "savings",
		Balance:     15000.00,
		Currency:    "USD",
		LastUpdated: time.Now().Add(-48 * time.Hour),
		IsActive:    true,
	},
	{
		ID:          "acc_3",
		Name:        "Credit Card",
		Bank:        "Wells Fargo",
		AccountType: "credit",
		Balance:     -1250.30,
		Currency:    "USD",
		LastUpdated: time.Now().Add(-12 * time.Hour),
		IsActive:    true,
	},
}

// Handler serves GET /api/accounts and GET /api/accounts?id=acc_1
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	id := r.URL.Query().Get("id")
	if id != "" {
		for _, a := range mockAccounts {
			if a.ID == id {
				writeJSON(w, http.StatusOK, APIResponse{Success: true, Message: "Account retrieved successfully", Data: a})
				return
			}
		}
		writeJSON(w, http.StatusNotFound, APIResponse{Success: false, Message: "Account not found"})
		return
	}

	if r.Method == http.MethodGet {
		writeJSON(w, http.StatusOK, APIResponse{Success: true, Message: "Accounts retrieved successfully", Data: mockAccounts})
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
