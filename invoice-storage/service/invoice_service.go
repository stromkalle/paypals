package service

import (
	"fmt"

	"example.com/invoice-storage/repository"
)

type InvoiceService interface {
	Save(string) bool
}

type invoiceService struct {
	invoiceRepository repository.InvoiceRepository
}

func NewInvoiceService(invoiceRepository repository.InvoiceRepository) *invoiceService {
	return &invoiceService{invoiceRepository: invoiceRepository}
}

func (s *invoiceService) Save(filepath string) bool {
	fmt.Println("Saving invoice in service: ", filepath)
	s.invoiceRepository.Save(filepath)
	return true
}
