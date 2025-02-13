package service

import (
	"encoding/csv"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"example.com/invoice-storage/domain"
	"example.com/invoice-storage/repository"
)

type InvoiceService interface {
	Save(multipart.File) (string, error)
	Get(string) (domain.Invoice, error)
}

type invoiceService struct {
	invoiceRepository repository.InvoiceRepository
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository) *invoiceService {
	return &invoiceService{invoiceRepository: invoiceRepository}
}

func (s *invoiceService) Save(file multipart.File) (string, error) {
	createInvoice, err := parse(file)

	if err != nil {
		return "", err
	}

	invoiceId, err := s.invoiceRepository.Save(createInvoice)
	if err != nil {
		return "", err
	}

	return invoiceId, nil
}

func (s *invoiceService) Get(invoiceId string) (domain.Invoice, error) {
	invoice, err := s.invoiceRepository.GetInvoice(invoiceId)
	if err != nil {
		return domain.Invoice{}, err
	}

	transactions, err := s.invoiceRepository.GetTransactions(invoiceId)
	if err != nil {
		return domain.Invoice{}, err
	}

	invoice.Transactions = transactions

	return invoice, nil
}

func parse(file multipart.File) (domain.CreateInvoice, error) {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return domain.CreateInvoice{}, err
	}

	transactions := []domain.CreateTransaction{}

	var minDate time.Time
	var maxDate time.Time

	for _, record := range records[1:] {
		amount, err := strconv.ParseFloat(strings.ReplaceAll(record[4], ",", "."), 64)
		if err != nil {
			log.Println("Unable to parse float: ", err)
			return domain.CreateInvoice{}, err
		}

		date, err := time.Parse("01/02/2006", record[0])
		if err != nil {
			log.Println("Unable to parse date: ", err)
			return domain.CreateInvoice{}, err
		}

		if minDate.IsZero() || date.Before(minDate) {
			minDate = date
		}

		if maxDate.IsZero() || date.After(maxDate) {
			maxDate = date
		}

		transactions = append(transactions, domain.CreateTransaction{
			Amount: amount,
			Date:   date,
			Buyer:  record[2],
			Vendor: record[1],
		})
	}

	return domain.CreateInvoice{
		StartDate:    minDate,
		EndDate:      maxDate,
		Transactions: transactions,
	}, nil

}
