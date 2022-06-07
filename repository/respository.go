package repository

import (
	"gemm123/pustaka-api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllBook() ([]models.Book, error)
	FindBookByID(ID int) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	DeleteBook(book models.Book) (models.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllBook() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindBookByID(ID int) (models.Book, error) {
	var book models.Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) UpdateBook(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) DeleteBook(book models.Book) (models.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}

func (r *repository) CreateBook(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}
