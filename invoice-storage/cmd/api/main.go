package main

import (
	"example.com/invoice-storage/database"
	"example.com/invoice-storage/handler"
	"example.com/invoice-storage/repository"
	"example.com/invoice-storage/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.New()
	invoiceRepository := repository.NewInvoiceRepository(db)
	invoiceService := service.NewInvoiceService(invoiceRepository)
	invoiceHandler := handler.NewInvoiceHandler(invoiceService)

	r := gin.Default()

	r.POST("/invoice", invoiceHandler.Save)
	r.GET("/invoice/:id", invoiceHandler.Get)

	r.Run()

}
