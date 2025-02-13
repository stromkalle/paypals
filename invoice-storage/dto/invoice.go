package dto

import "example.com/invoice-storage/domain"

type CreateInvoiceResponse struct {
	Invoice domain.Invoice `json:"invoice"`
}

type GetInvoiceRequest struct {
	ID string `json:"id"`
}

type GetInvoiceResponse struct {
	Invoice domain.Invoice `json:"invoice"`
}
