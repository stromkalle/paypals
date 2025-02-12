package handler

import (
	"fmt"
	"net/http"

	"example.com/invoice-storage/service"
	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	invoiceService service.InvoiceService
}

func NewInvoiceHandler(invoiceService service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: invoiceService}
}

func (h *InvoiceHandler) Save(c *gin.Context) {

	fmt.Println("Saving invoice in handler")
	h.invoiceService.Save("C://Users/hello-world")

	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
