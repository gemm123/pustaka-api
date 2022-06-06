package handler

import (
	"gemm123/pustaka-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"judul": "hello world!",
	})
}

func GetBooksByID(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetBooksQuery(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func PostBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": book.Title,
		"price": book.Price,
	})
}
