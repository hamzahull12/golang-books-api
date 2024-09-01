package main

import (
	"golang-books-api/handlers"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.POST("/books", handlers.AddBookHandler)
	router.Run(":9000")
}

