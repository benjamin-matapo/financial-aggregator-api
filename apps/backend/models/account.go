package models

import (
	"time"
)

// Account represents a bank account
type Account struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Bank        string    `json:"bank"`
	AccountType string    `json:"account_type"` // checking, savings, credit, investment
	Balance     float64   `json:"balance"`
	Currency    string    `json:"currency"`
	LastUpdated time.Time `json:"last_updated"`
	IsActive    bool      `json:"is_active"`
}

// AccountRefreshRequest represents a request to refresh account data
type AccountRefreshRequest struct {
	AccountID string `json:"account_id"`
}

// AccountRefreshResponse represents the response after refreshing account data
type AccountRefreshResponse struct {
	AccountID   string    `json:"account_id"`
	Success     bool      `json:"success"`
	Message     string    `json:"message"`
	LastUpdated time.Time `json:"last_updated"`
	NewBalance  float64   `json:"new_balance,omitempty"`
}
