package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"example.com/invoice-storage/domain"
)

type InvoiceRepository interface {
	Save(domain.CreateInvoice) (string, error)
	GetInvoice(string) (domain.Invoice, error)
	GetTransactions(string) ([]domain.Transaction, error)
}

type SQLInvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *SQLInvoiceRepository {
	return &SQLInvoiceRepository{db: db}
}

func (r *SQLInvoiceRepository) Save(invoice domain.CreateInvoice) (string, error) {
	fmt.Println("Saving invoice in repository")

	tx, err := r.db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	var invoiceID string
	err = tx.QueryRow("INSERT INTO invoices (start_date, end_date) VALUES ($1, $2) RETURNING id", invoice.StartDate, invoice.EndDate).Scan(&invoiceID)
	if err != nil {
		return "", fmt.Errorf("failed to insert invoice: %v", err)
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
		return "", fmt.Errorf("failed to bulk insert transactions: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	return invoiceID, nil
}

func (r *SQLInvoiceRepository) GetInvoice(id string) (domain.Invoice, error) {
	invoice := domain.Invoice{}

	row := r.db.QueryRow(
		"SELECT id, start_date, end_date FROM invoices WHERE id = $1",
		id,
	)
	err := row.Scan(
		&invoice.ID,
		&invoice.StartDate,
		&invoice.EndDate,
	)

	if err != nil {
		return domain.Invoice{}, fmt.Errorf("failed to query invoice: %v", err)
	}

	return invoice, nil
}

func (r *SQLInvoiceRepository) GetTransactions(id string) ([]domain.Transaction, error) {
	transactions := []domain.Transaction{}

	rows, err := r.db.Query(
		"SELECT id, amount, date, buyer, vendor FROM transactions WHERE invoice_id = $1",
		id,
	)

	if err != nil {
		return []domain.Transaction{}, fmt.Errorf("failed to query transactions: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var transaction domain.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.Amount,
			&transaction.Date,
			&transaction.Buyer,
			&transaction.Vendor,
		)

		if err != nil {
			return []domain.Transaction{}, fmt.Errorf("failed to query transactions: %v", err)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
