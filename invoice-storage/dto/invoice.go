package dto

import "example.com/invoice-storage/domain"

type CreateInvoiceResponse struct {
	ID string `json:"invoiceId"`
}
type GetInvoiceResponse struct {
	Invoice domain.Invoice `json:"invoice"`
}
