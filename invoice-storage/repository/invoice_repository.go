package repository

import (
	"database/sql"
	"fmt"
)

type InvoiceRepository interface {
	Save(string) bool
}

type SQLInvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *SQLInvoiceRepository {
	return &SQLInvoiceRepository{db: db}
}

func (r *SQLInvoiceRepository) Save(invoice string) bool {
	fmt.Println("Saving invoice in repository")
	return true
}
