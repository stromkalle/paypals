package main

import (
	"net/http"

	"exapmle.com/invoice-storage/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	repository.NewDatabase()

	r := gin.Default()

	r.POST("/invoice", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	r.Run()
}
