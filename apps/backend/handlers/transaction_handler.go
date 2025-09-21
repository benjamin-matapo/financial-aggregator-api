package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"financial-aggregator-api/backend/models"
	"financial-aggregator-api/backend/services"

	"github.com/go-chi/chi/v5"
)

// TransactionHandler handles transaction-related HTTP requests
type TransactionHandler struct {
	transactionService *services.TransactionService
}

// NewTransactionHandler creates a new TransactionHandler instance
func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

// GetTransactions handles GET /api/transactions
func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	filter := h.buildTransactionFilter(r)

	transactions, err := h.transactionService.GetAllTransactions(filter)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to fetch transactions", err)
		return
	}

	// Calculate pagination metadata
	total := len(transactions)
	limit := 50
	offset := 0

	if filter != nil {
		limit = filter.Limit
		offset = filter.Offset
	}

	if limit <= 0 {
		limit = 50
	}

	pages := (total + limit - 1) / limit

	response := models.PaginatedResponse{
		Success: true,
		Data:    transactions,
		Meta: models.PaginationMeta{
			Total:  total,
			Limit:  limit,
			Offset: offset,
			Pages:  pages,
		},
	}

	h.writeJSONResponse(w, http.StatusOK, response)
}

// GetTransactionByID handles GET /api/transactions/:id
func (h *TransactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	transactionID := chi.URLParam(r, "id")
	if transactionID == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "Transaction ID is required", nil)
		return
	}

	transaction, err := h.transactionService.GetTransactionByID(transactionID)
	if err != nil {
		h.writeErrorResponse(w, http.StatusNotFound, "Transaction not found", err)
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "Transaction retrieved successfully",
		Data:    transaction,
	}

	h.writeJSONResponse(w, http.StatusOK, response)
}

// GetTransactionsByAccount handles GET /api/accounts/:id/transactions
func (h *TransactionHandler) GetTransactionsByAccount(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "id")
	if accountID == "" {
		h.writeErrorResponse(w, http.StatusBadRequest, "Account ID is required", nil)
		return
	}

	// Parse limit from query parameter
	limit := 50
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	transactions, err := h.transactionService.GetTransactionsByAccountID(accountID, limit)
	if err != nil {
		h.writeErrorResponse(w, http.StatusInternalServerError, "Failed to fetch transactions", err)
		return
	}

	response := models.APIResponse{
		Success: true,
		Message: "Account transactions retrieved successfully",
		Data:    transactions,
	}

	h.writeJSONResponse(w, http.StatusOK, response)
}

// buildTransactionFilter builds a TransactionFilter from query parameters
func (h *TransactionHandler) buildTransactionFilter(r *http.Request) *models.TransactionFilter {
	filter := &models.TransactionFilter{}

	// Parse query parameters
	if accountID := r.URL.Query().Get("account_id"); accountID != "" {
		filter.AccountID = accountID
	}

	if transactionType := r.URL.Query().Get("type"); transactionType != "" {
		filter.Type = transactionType
	}

	if category := r.URL.Query().Get("category"); category != "" {
		filter.Category = category
	}

	if status := r.URL.Query().Get("status"); status != "" {
		filter.Status = status
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			filter.Limit = limit
		}
	}

	if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
		if offset, err := strconv.Atoi(offsetStr); err == nil && offset >= 0 {
			filter.Offset = offset
		}
	}

	if startDateStr := r.URL.Query().Get("start_date"); startDateStr != "" {
		if startDate, err := time.Parse("2006-01-02", startDateStr); err == nil {
			filter.StartDate = &startDate
		}
	}

	if endDateStr := r.URL.Query().Get("end_date"); endDateStr != "" {
		if endDate, err := time.Parse("2006-01-02", endDateStr); err == nil {
			filter.EndDate = &endDate
		}
	}

	return filter
}

// writeJSONResponse writes a JSON response to the client
func (h *TransactionHandler) writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// writeErrorResponse writes an error response to the client
func (h *TransactionHandler) writeErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	response := models.APIResponse{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	h.writeJSONResponse(w, statusCode, response)
}
