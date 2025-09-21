package models

import (
	"time"
)

// Transaction represents a financial transaction
type Transaction struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	Amount      float64   `json:"amount"`
	Currency    string    `json:"currency"`
	Type        string    `json:"type"`     // debit, credit, transfer
	Category    string    `json:"category"` // food, transportation, salary, etc.
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Status      string    `json:"status"` // pending, completed, failed, cancelled
	Reference   string    `json:"reference,omitempty"`
}

// TransactionFilter represents filters for querying transactions
type TransactionFilter struct {
	AccountID string     `json:"account_id,omitempty"`
	Type      string     `json:"type,omitempty"`
	Category  string     `json:"category,omitempty"`
	Status    string     `json:"status,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
	Limit     int        `json:"limit,omitempty"`
	Offset    int        `json:"offset,omitempty"`
}
