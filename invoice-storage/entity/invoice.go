package entity

import "time"

type InvoiceEntity struct {
	ID        string
	StartDate time.Time
	EndDate   time.Time
}
