package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"example.com/invoice-storage/domain"
)

type InvoiceRepository interface {
	Save(domain.CreateInvoice) (domain.Invoice, error)
	Get(int64) (domain.Invoice, error)
}

type SQLInvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *SQLInvoiceRepository {
	return &SQLInvoiceRepository{db: db}
}

func (r *SQLInvoiceRepository) Save(invoice domain.CreateInvoice) (domain.Invoice, error) {
	fmt.Println("Saving invoice in repository")

	tx, err := r.db.Begin()
	if err != nil {
		return domain.Invoice{}, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	var invoiceID int64
	err = tx.QueryRow("INSERT INTO invoices (start_date, end_date) VALUES ($1, $2) RETURNING id", invoice.StartDate, invoice.EndDate).Scan(&invoiceID)
	if err != nil {
		return domain.Invoice{}, fmt.Errorf("failed to insert invoice: %v", err)
	}

	valueStrings := make([]string, 0, len(invoice.Transactions))
	valueArgs := make([]interface{}, 0, len(invoice.Transactions)*5)

	for i, transaction := range invoice.Transactions {
		n := i*5 + 1
		placeholder := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", n, n+1, n+2, n+3, n+4)
		valueStrings = append(valueStrings, placeholder)
		valueArgs = append(valueArgs,
			invoiceID,
			transaction.Amount,
			transaction.Date,
			transaction.Buyer,
			transaction.Vendor,
		)
	}

	query := fmt.Sprintf(
		"INSERT INTO transactions (invoice_id, amount, date, buyer, vendor) VALUES %s",
		strings.Join(valueStrings, ","),
	)

	_, err = tx.Exec(query, valueArgs...)
	if err != nil {
		return domain.Invoice{}, fmt.Errorf("failed to bulk insert transactions: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return domain.Invoice{}, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return r.Get(invoiceID)
}

func (r *SQLInvoiceRepository) Get(id int64) (domain.Invoice, error) {
	invoice := domain.Invoice{}

	err := r.db.QueryRow("SELECT id, start_date, end_date FROM invoices WHERE id = $1", id).Scan(&invoice.ID, &invoice.StartDate, &invoice.EndDate)
	if err != nil {
		return domain.Invoice{}, fmt.Errorf("failed to query invoice: %v", err)
	}

	transactions, err := r.db.Query("SELECT id, amount, date, buyer, vendor FROM transactions WHERE invoice_id = $1", id)
	if err != nil {
		return domain.Invoice{}, fmt.Errorf("failed to query transactions: %v", err)
	}
	defer transactions.Close()

	for transactions.Next() {
		var transaction domain.Transaction
		err = transactions.Scan(&transaction.ID, &transaction.Amount, &transaction.Date, &transaction.Buyer, &transaction.Vendor)
		if err != nil {
			return domain.Invoice{}, fmt.Errorf("failed to scan transaction: %v", err)
		}
		invoice.Transactions = append(invoice.Transactions, transaction)
	}

	return invoice, nil
}
