package main

import (
	"example.com/invoice-storage/handler"
	"example.com/invoice-storage/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	repository.NewDatabase()
	ih := handler.NewInvoice()

	r := gin.Default()

	r.POST("/invoice", ih.StoreInvoice)
	r.Run()
}
