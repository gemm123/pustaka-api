package service

import (
	"gemm123/pustaka-api/models"
	"gemm123/pustaka-api/repository"
)

type Service interface {
	FindAllBook() ([]models.Book, error)
	FindBookByID(ID int) (models.Book, error)
	CreateBook(bookRequest models.Book) (models.Book, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAllBook() ([]models.Book, error) {
	books, err := s.repository.FindAllBook()
	return books, err
}

func (s *service) FindBookByID(ID int) (models.Book, error) {
	book, err := s.repository.FindBookByID(ID)
	return book, err
}

func (s *service) Create(bookRequest models.Book) (models.Book, error) {
	book := models.Book{
		Title: bookRequest.Title,
		Price: bookRequest.Price,
	}

	newBook, err := s.repository.CreateBook(book)

	return newBook, err
}
