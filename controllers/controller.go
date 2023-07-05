package controllers

import (
	"net/http"
	"strconv"
	"yuvraj/project/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := models.CreateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "book has been created"})
}

func GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	id2, _ := strconv.Atoi(id)

	book, err := models.GetBookByID(id2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if book == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	bookID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = models.DeleteBook(bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.UpdateBook(id, book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
func GetBooks(ctx *gin.Context) {
	books, err := models.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}
