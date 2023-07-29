package controllers

import (
	//"errors"
	"fmt"
	"net/http"
	"strconv"

	//"yuvraj/project/config"
	"yuvraj/models"

	"github.com/gin-gonic/gin"
)

func IssueBook(ctx *gin.Context) {
	var request models.IssuedBook
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	issuedBook, err := models.GetIssuedBookByBookIDAndUserID(strconv.Itoa(request.BookID), request.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue book"})
		fmt.Println(err)
		return
	}
	if issuedBook != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Book already issued to the user"})
		return
	}

	err = models.IssueBook(models.IssuedBook{
		UserID: request.UserID,
		BookID: request.BookID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to issue book"})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book issued successfully"})
}

func ReturnBook(ctx *gin.Context) {
	bookIDStr := ctx.Query("book_id")
	userIDStr := ctx.Query("user_id")

	if bookIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing book ID"})
		return
	}

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		fmt.Println(err)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		fmt.Println(err)
		return
	}

	err = models.ReturnBookFromUser(bookID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to return book"})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}
func ReissueBook(ctx *gin.Context) {
	var request models.IssuedBook
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := models.ReissueBook(request.BookID, request.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reissue book"})
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book reissued successfully"})
}
func CalculateFine(c *gin.Context) {
	bookIDStr := c.Query("book_id")
	userIDStr := c.Query("user_id")

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		fmt.Println(err)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	fine, err := models.CalculateFine(userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate fine"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"fine": fine})
}

func UpdateFine(c *gin.Context) {
	bookIDStr := c.Query("book_id")
	userIDStr := c.Query("user_id")

	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = models.UpdateFine(userID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update fine"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fine updated successfully"})
}
