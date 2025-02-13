package domain

import "time"

type CreateTransaction struct {
	Amount float64
	Date   time.Time
	Buyer  string
	Vendor string
}
type Transaction struct {
	ID     string
	Amount float64
	Date   time.Time
	Buyer  string
	Vendor string
}

type CreateInvoice struct {
	StartDate    time.Time
	EndDate      time.Time
	Transactions []CreateTransaction
}
type Invoice struct {
	ID           string
	StartDate    time.Time
	EndDate      time.Time
	Transactions []Transaction
}
