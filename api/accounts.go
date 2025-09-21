//go:build ignore
// +build ignore
package handler

import (
	"net/http"
	"../api/types"
)

// Mock data for demonstration
var mockAccounts = []types.Account{
	{
		ID:          "acc_1",
		Name:        "Main Checking",
		Bank:        "Chase Bank",
		AccountType: "checking",
		Balance:     2500.75,
		Currency:    "USD",
		LastUpdated: types.Now().Add(-24 * types.Hour),
		IsActive:    true,
	},
	{
		ID:          "acc_2",
		Name:        "Savings Account",
		Bank:        "Bank of America",
		AccountType: "savings",
		Balance:     15000.00,
		Currency:    "USD",
		LastUpdated: types.Now().Add(-48 * types.Hour),
		IsActive:    true,
	},
	{
		ID:          "acc_3",
		Name:        "Credit Card",
		Bank:        "Wells Fargo",
		AccountType: "credit",
		Balance:     -1250.30,
		Currency:    "USD",
		LastUpdated: types.Now().Add(-12 * types.Hour),
		IsActive:    true,
	},
}

// Handler is the Vercel serverless function for accounts
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

	// Get account ID from query parameter
	accountID := r.URL.Query().Get("id")

	if accountID != "" {
		// Get specific account
		for _, account := range mockAccounts {
			if account.ID == accountID {
				response := types.APIResponse{
					Success: true,
					Message: "Account retrieved successfully",
					Data:    account,
				}
				types.WriteJSONResponse(w, http.StatusOK, response)
				return
			}
		}
		
		// Account not found
		response := types.APIResponse{
			Success: false,
			Message: "Account not found",
		}
		types.WriteJSONResponse(w, http.StatusNotFound, response)
		return
	}

	// Get all accounts
	if r.Method == "GET" {
		response := types.APIResponse{
			Success: true,
			Message: "Accounts retrieved successfully",
			Data:    mockAccounts,
		}
		types.WriteJSONResponse(w, http.StatusOK, response)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}
