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
	Save(multipart.File) (domain.Invoice, error)
}

type invoiceService struct {
	invoiceRepository repository.InvoiceRepository
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository) *invoiceService {
	return &invoiceService{invoiceRepository: invoiceRepository}
}

func (s *invoiceService) Save(file multipart.File) (domain.Invoice, error) {
	createInvoice, err := parse(file)

	if err != nil {
		return domain.Invoice{}, err
	}

	savedInvoice, err := s.invoiceRepository.Save(createInvoice)
	if err != nil {
		return domain.Invoice{}, err
	}

	return savedInvoice, nil
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

// func parseCSV(csvFile multipart.File) ([]domain.Transaction, error) {
// 	csvReader := csv.NewReader(csvFile)
// 	records, err := csvReader.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 		transaction := Transaction{
// 			Date:       r[0],
// 			Vendor:     r[1],
// 			Transactor: r[2],
// 			Amount:     amount,
// 		}

// 	transactions := []domain.Transaction{}
// 	for _, record := range records[1:] {

// 		transactions = append(transactions, domain.Transaction{
// 			ID:        record[0],
// 			Amount:    record[1],
// 			Date:      record[2],
// 			Buyer:     record[3],
// 			Vendor:    record[4],
// 		})
// }
