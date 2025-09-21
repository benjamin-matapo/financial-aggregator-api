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

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Success bool           `json:"success"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Pages  int `json:"pages"`
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

// Handler is the Vercel serverless function for transactions
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

	// Get transaction ID from query parameter
	transactionID := r.URL.Query().Get("id")

	if transactionID != "" {
		// Get specific transaction
		for _, transaction := range mockTransactions {
			if transaction.ID == transactionID {
				response := APIResponse{
					Success: true,
					Message: "Transaction retrieved successfully",
					Data:    transaction,
				}
				writeJSONResponse(w, http.StatusOK, response)
				return
			}
		}
		
		// Transaction not found
		response := APIResponse{
			Success: false,
			Message: "Transaction not found",
		}
		writeJSONResponse(w, http.StatusNotFound, response)
		return
	}

	// Get all transactions with filtering
	if r.Method == "GET" {
		filteredTransactions := filterTransactions(r, mockTransactions)
		
		// Parse pagination parameters
		limit := 50
		offset := 0
		
		if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
			if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
				limit = parsedLimit
			}
		}
		
		if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
			if parsedOffset, err := strconv.Atoi(offsetStr); err == nil && parsedOffset >= 0 {
				offset = parsedOffset
			}
		}

		total := len(filteredTransactions)
		pages := (total + limit - 1) / limit

		// Apply pagination
		end := offset + limit
		if end > total {
			end = total
		}
		if offset > total {
			offset = total
		}

		paginatedTransactions := filteredTransactions[offset:end]

		response := PaginatedResponse{
			Success: true,
			Data:    paginatedTransactions,
			Meta: PaginationMeta{
				Total:  total,
				Limit:  limit,
				Offset: offset,
				Pages:  pages,
			},
		}

		writeJSONResponse(w, http.StatusOK, response)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func filterTransactions(r *http.Request, transactions []Transaction) []Transaction {
	var filtered []Transaction
	
	accountID := r.URL.Query().Get("account_id")
	transactionType := r.URL.Query().Get("type")
	category := r.URL.Query().Get("category")
	status := r.URL.Query().Get("status")

	for _, txn := range transactions {
		if accountID != "" && txn.AccountID != accountID {
			continue
		}
		if transactionType != "" && txn.Type != transactionType {
			continue
		}
		if category != "" && txn.Category != category {
			continue
		}
		if status != "" && txn.Status != status {
			continue
		}
		filtered = append(filtered, txn)
	}

	return filtered
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
