package handlers

import (
	"fmt"
	"golang-books-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var books []models.Book

func AddBookHandler(ctx *gin.Context) {
	var book models.Book
	if err := ctx.BindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "fail", "message": err.Error()})
		return
	}

	if book.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Gagal menambahkan buku. Mohon isi nama buku",
		})
		return
	}
	
	if book.ReadPage > book.PageCount {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
		})
		return
	}
	
	book.ID = fmt.Sprintf("note-%s", gonanoid.MustGenerate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", 16))
	book.Finished = book.PageCount == book.ReadPage
	book.InsertedAt = time.Now().UTC()
	book.UpdatedAt = book.InsertedAt

	books = append(books, book)

	response := models.CreateBookResponse{
		Status: "success",
		Message: "Buku berhasil ditambahkan",
	}
	response.Data.BookID = book.ID
	ctx.JSON(http.StatusCreated, response)
}

func GetBooksAllHandler(ctx *gin.Context) {
	bookSummary := []models.BookSummary{}

	for _, book := range books{
		bookSummary = append(bookSummary, models.BookSummary{
			ID: book.ID,
			Name: book.Name,
			Publisher: book.Publisher,
		})
	}

	response := models.GetAllBooksResponse{
		Status: "success",
	}
	response.Data.Books = bookSummary
	ctx.JSON(http.StatusOK, response)
}

func GetBookByIdHandler(ctx *gin.Context) {
	bookID := ctx.Param("id")
	var foundBook *models.Book

	for _, book := range books {
		if book.ID == bookID {
			foundBook = &book
			break
		}
	}

	if foundBook != nil {
		response := models.GetBookByIdResponse{
			Status: "success",
		}
		response.Data.Book = foundBook
		ctx.JSON(http.StatusOK, response)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "Buku tidak ditemukan",
		})
	}
}

func EditBookByIdHandler(ctx *gin.Context) {
	bookId := ctx.Param("id")
	var updateRequest models.Book

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"message": "Permintaan tidak valid",
		})
		return
	}

	var foundBook *models.Book
	for i, book := range books {
		if book.ID == bookId {
			foundBook = &books[i]
			break
		}
	}

	if foundBook == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": "fail",
			"message": "Gagal memperbarui buku. Id tidak ditemukan",
		})
		return
	}
	if updateRequest.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui buku. Mohon isi nama buku",
		})
		return
	}
	if updateRequest.ReadPage > updateRequest.PageCount {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount",
		})
		return
	}

	foundBook.Name = updateRequest.Name
	foundBook.Year = updateRequest.Year
	foundBook.Author = updateRequest.Author
	foundBook.Summary = updateRequest.Summary
	foundBook.Publisher = updateRequest.Publisher
	foundBook.PageCount = updateRequest.PageCount
	foundBook.ReadPage = updateRequest.ReadPage
	foundBook.Reading = updateRequest.Reading
	foundBook.Finished = updateRequest.PageCount == updateRequest.ReadPage
	foundBook.UpdatedAt = time.Now().UTC()
	
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Buku berhasil diperbarui",
	})
}