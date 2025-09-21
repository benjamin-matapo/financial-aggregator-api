package handlers

import (
	"encoding/json"
	"net/http"

	"financial-aggregator-api/backend/models"
	"financial-aggregator-api/backend/services"

	"github.com/go-chi/chi/v5"
)

// AccountHandler handles account-related HTTP requests
type AccountHandler struct {
	accountService *services.AccountService
}

// NewAccountHandler creates a new AccountHandler instance
func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

// GetAccounts handles GET /api/accounts
func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accountService.GetAllAccounts()
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to fetch accounts", err)
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "Accounts retrieved successfully",
		Data:    accounts,
	}

	h.writeJSONResponse(w, http.StatusOK, response)
}

// GetAccountByID handles GET /api/accounts/:id
func (h *AccountHandler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "id")
	if accountID == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "Account ID is required", nil)
		return
	}

	account, err := h.accountService.GetAccountByID(accountID)
	if err != nil {
		h.writeErrorResponse(w, http.StatusNotFound, "Account not found", err)
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "Account retrieved successfully",
		Data:    account,
	}

	h.writeJSONResponse(w, http.StatusOK, response)
}

// RefreshAccount handles POST /api/accounts/:id/refresh
func (h *AccountHandler) RefreshAccount(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "id")
	if accountID == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "Account ID is required", nil)
		return
	}

	refreshResponse, err := h.accountService.RefreshAccount(accountID)
	if err != nil {
		h.writeErrorResponse(w, http.StatusNotFound, "Account not found", err)
		return
	}

	statusCode := http.StatusOK
	if !refreshResponse.Success {
		statusCode = http.StatusBadRequest
	}

	response := models.APIResponse{
		Success: refreshResponse.Success,
		Message: refreshResponse.Message,
		Data:    refreshResponse,
	}

	h.writeJSONResponse(w, statusCode, response)
}

// writeJSONResponse writes a JSON response to the client
func (h *AccountHandler) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// writeErrorResponse writes an error response to the client
func (h *AccountHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	response := models.APIResponse{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	h.writeJSONResponse(w, statusCode, response)
}
