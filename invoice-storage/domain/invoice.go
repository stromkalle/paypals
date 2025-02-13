package domain

import "time"

type Transaction struct {
	ID     string
	Amount float64
	Date   time.Time
	Buyer  string
	Vendor string
}

type Invoice struct {
	ID           string
	Filepath     string
	StartDate    time.Time
	EndDate      time.Time
	Transactions []Transaction
}
