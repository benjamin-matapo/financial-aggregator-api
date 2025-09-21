package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"financial-aggregator-api/backend/models"
	"financial-aggregator-api/backend/services"

	"github.com/go-chi/chi/v5"
)

func TestTransactionHandler_GetTransactions(t *testing.T) {
	// Create mock service
	transactionService := services.NewTransactionService()
	handler := NewTransactionHandler(transactionService)

	// Test get all transactions
	req, err := http.NewRequest("GET", "/api/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/transactions", handler.GetTransactions)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.PaginatedResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Errorf("Expected success to be true, got %v", response.Success)
	}

	// Check that we have transactions
	transactions, ok := response.Data.([]interface{})
	if !ok {
		t.Fatal("Expected transactions array in response data")
	}

	if len(transactions) == 0 {
		t.Error("Expected at least one transaction in response")
	}
}

func TestTransactionHandler_GetTransactionsWithFilters(t *testing.T) {
	// Create mock service
	transactionService := services.NewTransactionService()
	handler := NewTransactionHandler(transactionService)

	// Test with account filter
	req, err := http.NewRequest("GET", "/api/transactions?account_id=acc_001", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/transactions", handler.GetTransactions)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.PaginatedResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Errorf("Expected success to be true, got %v", response.Success)
	}

	// Test with type filter
	req, err = http.NewRequest("GET", "/api/transactions?type=debit", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Test with limit
	req, err = http.NewRequest("GET", "/api/transactions?limit=5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var limitedResponse models.PaginatedResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &limitedResponse); err != nil {
		t.Fatal(err)
	}

	if limitedResponse.Meta.Limit != 5 {
		t.Errorf("Expected limit to be 5, got %v", limitedResponse.Meta.Limit)
	}
}

func TestTransactionHandler_GetTransactionByID(t *testing.T) {
	// Create mock service
	transactionService := services.NewTransactionService()
	handler := NewTransactionHandler(transactionService)

	// Test valid transaction ID
	req, err := http.NewRequest("GET", "/api/transactions/txn_001", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/transactions/{id}", handler.GetTransactionByID)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.APIResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Errorf("Expected success to be true, got %v", response.Success)
	}

	// Test invalid transaction ID
	req, err = http.NewRequest("GET", "/api/transactions/invalid_id", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestTransactionHandler_GetTransactionsByAccount(t *testing.T) {
	// Create mock service
	transactionService := services.NewTransactionService()
	handler := NewTransactionHandler(transactionService)

	// Test valid account ID
	req, err := http.NewRequest("GET", "/api/accounts/acc_001/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/accounts/{id}/transactions", handler.GetTransactionsByAccount)

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.APIResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Errorf("Expected success to be true, got %v", response.Success)
	}

	// Test with limit parameter
	req, err = http.NewRequest("GET", "/api/accounts/acc_001/transactions?limit=3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Test invalid account ID
	req, err = http.NewRequest("GET", "/api/accounts/invalid_id/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Should still return 200 with empty array for invalid account
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
