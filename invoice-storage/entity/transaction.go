package entity

import "time"

type TransactionEntity struct {
	ID        string
	InvoiceID string // Foreign key
	Amount    float64
	Date      time.Time
	Buyer     string
	Vendor    string
}
