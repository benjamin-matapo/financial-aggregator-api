//go:build ignore
// +build ignore
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Transaction represents a transaction
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

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Mock transactions data
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
}

// Handler is the Vercel serverless function for account transactions
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get account ID from query parameter
	accountID := r.URL.Query().Get("id")
	if accountID == "" {
		response := APIResponse{
			Success: false,
			Message: "Account ID is required",
		}
		writeJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// Parse limit from query parameter
	limit := 50
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	// Filter transactions by account ID
	var accountTransactions []Transaction
	for _, txn := range mockTransactions {
		if txn.AccountID == accountID {
			accountTransactions = append(accountTransactions, txn)
		}
	}

	// Apply limit
	if len(accountTransactions) > limit {
		accountTransactions = accountTransactions[:limit]
	}

	response := APIResponse{
		Success: true,
		Message: "Account transactions retrieved successfully",
		Data:    accountTransactions,
	}

	writeJSONResponse(w, http.StatusOK, response)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
