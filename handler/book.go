package handler

import (
	"gemm123/pustaka-api/models"
	"gemm123/pustaka-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service}
}

func (h *handler) GetAllBooks(c *gin.Context) {
	books, err := h.service.FindAllBook()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bookResponses []models.BookResponse

	for _, b := range books {
		bookResponse := models.BookResponse{
			ID:    int(b.ID),
			Title: b.Title,
			Price: b.Price,
		}

		bookResponses = append(bookResponses, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponses,
	})

}

func (h *handler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	book, err := h.service.FindBookByID(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookResponse := models.BookResponse{
		ID:    int(book.ID),
		Title: book.Title,
		Price: book.Price,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *handler) UpdateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	bookRes, err := h.service.UpdateBook(idInt, newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookResponse := models.BookResponse{
		Title: bookRes.Title,
		Price: bookRes.Price,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"title": bookResponse.Title,
			"price": bookResponse.Price,
		},
	})
}

func (h *handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	bookRes, err := h.service.DeleteBook(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookResponse := models.BookResponse{
		Title: bookRes.Title,
		Price: bookRes.Price,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"title": bookResponse.Title,
			"price": bookResponse.Price,
		},
	})
}

func (h *handler) PostBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook, err := h.service.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": newBook,
	})
}
