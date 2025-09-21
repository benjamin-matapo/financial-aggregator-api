//go:build ignore
// +build ignore
package handler

import (
    "encoding/json"
    "net/http"
    "time"
)

// Shared types for all handlers

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Shared JSON response helper used by all handlers
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

type PaginationMeta struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Pages  int `json:"pages"`
}

type PaginatedResponse struct {
	Success bool           `json:"success"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
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

type Transaction struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Type        string    `json:"type"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status"`
	Reference   string    `json:"reference,omitempty"`
}

// Mock data for demo
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

var mockTransactions = []Transaction{
	{
		ID:          "txn_1",
		AccountID:   "acc_1",
		Amount:      -45.67,
		Currency:    "USD",
		Type:        "debit",
		Category:    "groceries",
		Description: "Whole Foods Market",
		Date:        time.Now().Add(-24 * time.Hour),
		Status:      "completed",
		Reference:   "REF001",
	},
	{
		ID:          "txn_2",
		AccountID:   "acc_1",
		Amount:      2500.00,
		Currency:    "USD",
		Type:        "credit",
		Category:    "salary",
		Description: "Monthly Salary",
		Date:        time.Now().Add(-72 * time.Hour),
		Status:      "completed",
		Reference:   "SAL001",
	},
	{
		ID:          "txn_3",
		AccountID:   "acc_2",
		Amount:      -200.00,
		Currency:    "USD",
		Type:        "transfer",
		Category:    "transfer",
		Description: "Transfer to Checking",
		Date:        time.Now().Add(-48 * time.Hour),
		Status:      "completed",
		Reference:   "TRF001",
	},
	{
		ID:          "txn_4",
		AccountID:   "acc_3",
		Amount:      -89.99,
		Currency:    "USD",
		Type:        "debit",
		Category:    "shopping",
		Description: "Amazon Purchase",
		Date:        time.Now().Add(-36 * time.Hour),
		Status:      "completed",
		Reference:   "AMZ001",
	},
}
