package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
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

var mockTransactions = []Transaction{
	{ID: "txn_1", AccountID: "acc_1", Amount: -45.67, Currency: "USD", Type: "debit", Category: "groceries", Description: "Whole Foods Market", Date: time.Now().Add(-24 * time.Hour), Status: "completed", Reference: "REF001"},
	{ID: "txn_2", AccountID: "acc_1", Amount: 2500.00, Currency: "USD", Type: "credit", Category: "salary", Description: "Monthly Salary", Date: time.Now().Add(-72 * time.Hour), Status: "completed", Reference: "SAL001"},
	{ID: "txn_3", AccountID: "acc_2", Amount: -200.00, Currency: "USD", Type: "transfer", Category: "transfer", Description: "Transfer to Checking", Date: time.Now().Add(-48 * time.Hour), Status: "completed", Reference: "TRF001"},
}

// Handler serves GET /api/accounts/transactions?id=acc_1&limit=50
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[accounts/transactions]", r.Method, r.URL.String())
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	accountID := r.URL.Query().Get("id")
	if accountID == "" {
		writeJSON(w, http.StatusBadRequest, APIResponse{Success: false, Message: "Account ID is required"})
		return
	}

	limit := 50
	if s := r.URL.Query().Get("limit"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 {
			limit = v
		}
	}

	var items []Transaction
	for _, t := range mockTransactions {
		if t.AccountID == accountID {
			items = append(items, t)
		}
	}
	if len(items) > limit {
		items = items[:limit]
	}

	writeJSON(w, http.StatusOK, APIResponse{Success: true, Message: "Account transactions retrieved successfully", Data: items})
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
