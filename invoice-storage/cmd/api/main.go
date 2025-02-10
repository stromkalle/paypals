package main

import (
	"example.com/invoice-storage/handler"
	"example.com/invoice-storage/repository"
	"example.com/invoice-storage/service"
	"github.com/gin-gonic/gin"
)

func main() {

	repository.NewDatabase()
	handler.NewInvoiceHandler()
	service.NewInvoiceService()

	r := gin.Default()

	r.POST("/invoice", ih.StoreInvoice)
	r.Run()
}
