package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"financial-aggregator-api/backend/models"
	"financial-aggregator-api/backend/services"

	"github.com/go-chi/chi/v5"
)

func TestAccountHandler_GetAccounts(t *testing.T) {
	// Create mock service
	accountService := services.NewAccountService()
	handler := NewAccountHandler(accountService)

	// Create request
	req, err := http.NewRequest("GET", "/api/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create response recorder
	rr := httptest.NewRecorder()

	// Create chi router and add route
	r := chi.NewRouter()
	r.Get("/api/accounts", handler.GetAccounts)

	// Perform request
	r.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var response models.APIResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Errorf("Expected success to be true, got %v", response.Success)
	}

	// Check that we have accounts
	accounts, ok := response.Data.([]interface{})
	if !ok {
		t.Fatal("Expected accounts array in response data")
	}

	if len(accounts) == 0 {
		t.Error("Expected at least one account in response")
	}
}

func TestAccountHandler_GetAccountByID(t *testing.T) {
	// Create mock service
	accountService := services.NewAccountService()
	handler := NewAccountHandler(accountService)

	// Test valid account ID
	req, err := http.NewRequest("GET", "/api/accounts/acc_001", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/accounts/{id}", handler.GetAccountByID)

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

	// Test invalid account ID
	req, err = http.NewRequest("GET", "/api/accounts/invalid_id", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestAccountHandler_RefreshAccount(t *testing.T) {
	// Create mock service
	accountService := services.NewAccountService()
	handler := NewAccountHandler(accountService)

	// Test valid account refresh
	req, err := http.NewRequest("POST", "/api/accounts/acc_001/refresh", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Post("/api/accounts/{id}/refresh", handler.RefreshAccount)

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

	// Test invalid account refresh
	req, err = http.NewRequest("POST", "/api/accounts/invalid_id/refresh", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestAccountHandler_RefreshAccountWithBody(t *testing.T) {
	// Create mock service
	accountService := services.NewAccountService()
	handler := NewAccountHandler(accountService)

	// Test refresh with request body
	refreshRequest := models.AccountRefreshRequest{
		AccountID: "acc_001",
	}

	jsonData, err := json.Marshal(refreshRequest)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/accounts/acc_001/refresh", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Post("/api/accounts/{id}/refresh", handler.RefreshAccount)

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
}
