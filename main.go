package main

import (
	"golang-books-api/handlers"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.POST("/books", handlers.AddBookHandler)
	router.GET("/books", handlers.GetBooksAllHandler)
	router.GET("/books/:id", handlers.GetBookByIdHandler)
	router.Run(":9000")
}

