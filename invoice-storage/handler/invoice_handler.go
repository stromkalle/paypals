package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct{}

func NewInvoice() *InvoiceHandler {
	return &InvoiceHandler{}
}

func (h *InvoiceHandler) StoreInvoice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
