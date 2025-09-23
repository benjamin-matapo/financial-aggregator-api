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
	{ID: "txn_4", AccountID: "acc_3", Amount: -89.99, Currency: "USD", Type: "debit", Category: "shopping", Description: "Amazon Purchase", Date: time.Now().Add(-36 * time.Hour), Status: "completed", Reference: "AMZ001"},
}

// Handler serves GET /api/transactions and GET /api/transactions?id=txn_1
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[transactions]", r.Method, r.URL.String())
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	id := r.URL.Query().Get("id")
	if id != "" {
		for _, t := range mockTransactions {
			if t.ID == id {
				writeJSON(w, http.StatusOK, APIResponse{Success: true, Message: "Transaction retrieved successfully", Data: t})
				return
			}
		}
		writeJSON(w, http.StatusNotFound, APIResponse{Success: false, Message: "Transaction not found"})
		return
	}

	if r.Method == http.MethodGet {
		items := filter(r)
		limit := 50
		offset := 0
		if s := r.URL.Query().Get("limit"); s != "" {
			if v, err := strconv.Atoi(s); err == nil && v > 0 { limit = v }
		}
		if s := r.URL.Query().Get("offset"); s != "" {
			if v, err := strconv.Atoi(s); err == nil && v >= 0 { offset = v }
		}
		total := len(items)
		pages := (total + limit - 1) / limit
		end := offset + limit
		if end > total { end = total }
		if offset > total { offset = total }
		page := items[offset:end]
		writeJSON(w, http.StatusOK, PaginatedResponse{Success: true, Data: page, Meta: PaginationMeta{Total: total, Limit: limit, Offset: offset, Pages: pages}})
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func filter(r *http.Request) []Transaction {
	var out []Transaction
	acc := r.URL.Query().Get("account_id")
	typ := r.URL.Query().Get("type")
	cat := r.URL.Query().Get("category")
	st := r.URL.Query().Get("status")
	for _, t := range mockTransactions {
		if acc != "" && t.AccountID != acc { continue }
		if typ != "" && t.Type != typ { continue }
		if cat != "" && t.Category != cat { continue }
		if st != "" && t.Status != st { continue }
		out = append(out, t)
	}
	return out
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
