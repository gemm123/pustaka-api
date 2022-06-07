package service

import (
	"gemm123/pustaka-api/models"
	"gemm123/pustaka-api/repository"
)

type Service interface {
	FindAllBook() ([]models.Book, error)
	FindBookByID(ID int) (models.Book, error)
	UpdateBook(ID int, bookRequest models.Book) (models.Book, error)
	CreateBook(bookRequest models.Book) (models.Book, error)
	DeleteBook(ID int) (models.Book, error)
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

func (s *service) UpdateBook(ID int, bookRequest models.Book) (models.Book, error) {
	book, err := s.repository.FindBookByID(ID)

	book.Title = bookRequest.Title
	book.Price = bookRequest.Price

	newBook, err := s.repository.UpdateBook(book)

	return newBook, err
}

func (s *service) DeleteBook(ID int) (models.Book, error) {
	book, err := s.repository.FindBookByID(ID)

	bookResponse, err := s.repository.DeleteBook(book)

	return bookResponse, err
}

func (s *service) CreateBook(bookRequest models.Book) (models.Book, error) {
	book := models.Book{
		Title: bookRequest.Title,
		Price: bookRequest.Price,
	}

	newBook, err := s.repository.CreateBook(book)

	return newBook, err
}
