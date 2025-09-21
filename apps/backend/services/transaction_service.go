package services

import (
	"errors"
	"sort"
	"sync"
	"time"

	"financial-aggregator-api/backend/models"
)

// TransactionService handles transaction-related business logic
type TransactionService struct {
	transactions map[string]*models.Transaction
	mutex        sync.RWMutex
}

// NewTransactionService creates a new TransactionService instance
func NewTransactionService() *TransactionService {
	service := &TransactionService{
		transactions: make(map[string]*models.Transaction),
	}
	service.initializeMockData()
	return service
}

// GetAllTransactions returns all transactions with optional filtering
func (s *TransactionService) GetAllTransactions(filter *models.TransactionFilter) ([]*models.Transaction, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	var transactions []*models.Transaction

	// Convert map to slice
	for _, transaction := range s.transactions {
		transactions = append(transactions, transaction)
	}

	// Apply filters
	transactions = s.applyFilters(transactions, filter)

	// Sort by date (newest first)
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date.After(transactions[j].Date)
	})

	// Apply pagination
	if filter != nil {
		offset := filter.Offset
		limit := filter.Limit

		if limit <= 0 {
			limit = 50 // default limit
		}

		start := offset
		end := offset + limit

		if start >= len(transactions) {
			return []*models.Transaction{}, nil
		}

		if end > len(transactions) {
			end = len(transactions)
		}

		transactions = transactions[start:end]
	}

	return transactions, nil
}

// GetTransactionByID returns a transaction by ID
func (s *TransactionService) GetTransactionByID(id string) (*models.Transaction, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	transaction, exists := s.transactions[id]
	if !exists {
		return nil, errors.New("transaction not found")
	}

	return transaction, nil
}

// GetTransactionsByAccountID returns transactions for a specific account
func (s *TransactionService) GetTransactionsByAccountID(accountID string, limit int) ([]*models.Transaction, error) {
	filter := &models.TransactionFilter{
		AccountID: accountID,
		Limit:     limit,
	}
	return s.GetAllTransactions(filter)
}

// applyFilters applies the given filters to the transactions
func (s *TransactionService) applyFilters(transactions []*models.Transaction, filter *models.TransactionFilter) []*models.Transaction {
	if filter == nil {
		return transactions
	}

	var filtered []*models.Transaction

	for _, transaction := range transactions {
		// Account ID filter
		if filter.AccountID != "" && transaction.AccountID != filter.AccountID {
			continue
		}

		// Type filter
		if filter.Type != "" && transaction.Type != filter.Type {
			continue
		}

		// Category filter
		if filter.Category != "" && transaction.Category != filter.Category {
			continue
		}

		// Status filter
		if filter.Status != "" && transaction.Status != filter.Status {
			continue
		}

		// Date range filter
		if filter.StartDate != nil && transaction.Date.Before(*filter.StartDate) {
			continue
		}

		if filter.EndDate != nil && transaction.Date.After(*filter.EndDate) {
			continue
		}

		filtered = append(filtered, transaction)
	}

	return filtered
}

// initializeMockData populates the service with mock data
func (s *TransactionService) initializeMockData() {
	now := time.Now()

	mockTransactions := []*models.Transaction{
		{
			ID:          "txn_001",
			AccountID:   "acc_001",
			Amount:      -45.50,
			Currency:    "USD",
			Type:        "debit",
			Category:    "food",
			Description: "Grocery Store Purchase",
			Date:        now.Add(-2 * time.Hour),
			Status:      "completed",
			Reference:   "TXN001234567",
		},
		{
			ID:          "txn_002",
			AccountID:   "acc_001",
			Amount:      5000.00,
			Currency:    "USD",
			Type:        "credit",
			Category:    "salary",
			Description: "Monthly Salary",
			Date:        now.Add(-1 * time.Hour),
			Status:      "completed",
			Reference:   "SAL001234567",
		},
		{
			ID:          "txn_003",
			AccountID:   "acc_003",
			Amount:      -120.00,
			Currency:    "USD",
			Type:        "debit",
			Category:    "utilities",
			Description: "Electric Bill",
			Date:        now.Add(-3 * time.Hour),
			Status:      "completed",
			Reference:   "UTL001234567",
		},
		{
			ID:          "txn_004",
			AccountID:   "acc_002",
			Amount:      500.00,
			Currency:    "USD",
			Type:        "credit",
			Category:    "transfer",
			Description: "Transfer from Checking",
			Date:        now.Add(-4 * time.Hour),
			Status:      "completed",
			Reference:   "TRF001234567",
		},
		{
			ID:          "txn_005",
			AccountID:   "acc_001",
			Amount:      -25.00,
			Currency:    "USD",
			Type:        "debit",
			Category:    "transportation",
			Description: "Gas Station",
			Date:        now.Add(-5 * time.Hour),
			Status:      "completed",
			Reference:   "GAS001234567",
		},
		{
			ID:          "txn_006",
			AccountID:   "acc_004",
			Amount:      150.00,
			Currency:    "USD",
			Type:        "credit",
			Category:    "investment",
			Description: "Dividend Payment",
			Date:        now.Add(-6 * time.Hour),
			Status:      "completed",
			Reference:   "DIV001234567",
		},
		{
			ID:          "txn_007",
			AccountID:   "acc_001",
			Amount:      -80.00,
			Currency:    "USD",
			Type:        "debit",
			Category:    "entertainment",
			Description: "Movie Theater",
			Date:        now.Add(-24 * time.Hour),
			Status:      "completed",
			Reference:   "ENT001234567",
		},
		{
			ID:          "txn_008",
			AccountID:   "acc_005",
			Amount:      2500.00,
			Currency:    "USD",
			Type:        "credit",
			Category:    "business",
			Description: "Client Payment",
			Date:        now.Add(-48 * time.Hour),
			Status:      "completed",
			Reference:   "BIZ001234567",
		},
		{
			ID:          "txn_009",
			AccountID:   "acc_001",
			Amount:      -200.00,
			Currency:    "USD",
			Type:        "debit",
			Category:    "healthcare",
			Description: "Doctor Visit",
			Date:        now.Add(-72 * time.Hour),
			Status:      "completed",
			Reference:   "HLT001234567",
		},
		{
			ID:          "txn_010",
			AccountID:   "acc_002",
			Amount:      1000.00,
			Currency:    "USD",
			Type:        "credit",
			Category:    "transfer",
			Description: "Emergency Fund Contribution",
			Date:        now.Add(-96 * time.Hour),
			Status:      "completed",
			Reference:   "EMG001234567",
		},
	}

	for _, transaction := range mockTransactions {
		s.transactions[transaction.ID] = transaction
	}
}
