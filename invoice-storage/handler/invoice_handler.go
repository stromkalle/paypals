package handler

import (
	"fmt"
	"net/http"

	"example.com/invoice-storage/dto"
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

	fh, err := c.FormFile("csv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get CSV file",
		})
		return
	}

	csvFile, err := fh.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to open CSV file",
		})
		return
	}

	defer csvFile.Close()

	invoiceId, err := h.invoiceService.Save(csvFile)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to save invoice",
		})
		return
	}

	c.JSON(http.StatusBadRequest, dto.CreateInvoiceResponse{
		ID: invoiceId,
	})
}

func (h *InvoiceHandler) Get(c *gin.Context) {
	id := c.Param("id")

	invoice, err := h.invoiceService.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get invoice",
		})
		return
	}

	c.JSON(http.StatusBadRequest, dto.GetInvoiceResponse{
		Invoice: invoice,
	})
}
