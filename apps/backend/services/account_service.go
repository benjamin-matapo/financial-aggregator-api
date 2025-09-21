package services

import (
	"errors"
	"sync"
	"time"

	"financial-aggregator-api/backend/models"
)

// AccountService handles account-related business logic
type AccountService struct {
	accounts map[string]*models.Account
	mutex    sync.RWMutex
}

// NewAccountService creates a new AccountService instance
func NewAccountService() *AccountService {
	service := &AccountService{
		accounts: make(map[string]*models.Account),
	}
	service.initializeMockData()
	return service
}

// GetAllAccounts returns all accounts
func (s *AccountService) GetAllAccounts() ([]*models.Account, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	accounts := make([]*models.Account, 0, len(s.accounts))
	for _, account := range s.accounts {
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// GetAccountByID returns an account by ID
func (s *AccountService) GetAccountByID(id string) (*models.Account, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	account, exists := s.accounts[id]
	if !exists {
		return nil, errors.New("account not found")
	}

	return account, nil
}

// RefreshAccount simulates fetching updated data from external sources
func (s *AccountService) RefreshAccount(accountID string) (*models.AccountRefreshResponse, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	account, exists := s.accounts[accountID]
	if !exists {
		return &models.AccountRefreshResponse{
			AccountID:   accountID,
			Success:     false,
			Message:     "account not found",
			LastUpdated: time.Now(),
		}, errors.New("account not found")
	}

	// Simulate external API call delay
	time.Sleep(100 * time.Millisecond)

	// Simulate balance update (random change between -100 and +100)
	balanceChange := (float64(time.Now().UnixNano()%200) - 100) / 100
	account.Balance += balanceChange
	account.LastUpdated = time.Now()

	// Ensure balance doesn't go negative for checking/savings accounts
	if account.AccountType == "checking" || account.AccountType == "savings" {
		if account.Balance < 0 {
			account.Balance = 0
		}
	}

	return &models.AccountRefreshResponse{
		AccountID:   accountID,
		Success:     true,
		Message:     "account data refreshed successfully",
		LastUpdated: account.LastUpdated,
		NewBalance:  account.Balance,
	}, nil
}

// initializeMockData populates the service with mock data
func (s *AccountService) initializeMockData() {
	now := time.Now()

	mockAccounts := []*models.Account{
		{
			ID:          "acc_001",
			Name:        "Primary Checking",
			Bank:        "Chase Bank",
			AccountType: "checking",
			Balance:     2500.75,
			Currency:    "USD",
			LastUpdated: now.Add(-2 * time.Hour),
			IsActive:    true,
		},
		{
			ID:          "acc_002",
			Name:        "High Yield Savings",
			Bank:        "Ally Bank",
			AccountType: "savings",
			Balance:     15000.00,
			Currency:    "USD",
			LastUpdated: now.Add(-1 * time.Hour),
			IsActive:    true,
		},
		{
			ID:          "acc_003",
			Name:        "Credit Card",
			Bank:        "Capital One",
			AccountType: "credit",
			Balance:     -1200.50,
			Currency:    "USD",
			LastUpdated: now.Add(-30 * time.Minute),
			IsActive:    true,
		},
		{
			ID:          "acc_004",
			Name:        "Investment Account",
			Bank:        "Fidelity",
			AccountType: "investment",
			Balance:     45000.25,
			Currency:    "USD",
			LastUpdated: now.Add(-15 * time.Minute),
			IsActive:    true,
		},
		{
			ID:          "acc_005",
			Name:        "Business Checking",
			Bank:        "Wells Fargo",
			AccountType: "checking",
			Balance:     8500.00,
			Currency:    "USD",
			LastUpdated: now.Add(-45 * time.Minute),
			IsActive:    true,
		},
		{
			ID:          "acc_006",
			Name:        "Emergency Fund",
			Bank:        "Marcus by Goldman Sachs",
			AccountType: "savings",
			Balance:     25000.00,
			Currency:    "USD",
			LastUpdated: now.Add(-1 * time.Hour),
			IsActive:    true,
		},
	}

	for _, account := range mockAccounts {
		s.accounts[account.ID] = account
	}
}
