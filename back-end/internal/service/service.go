package service

import (
	"github.com/ZeeeUs/SimpleApp/internal/models"

	"github.com/rs/zerolog"
)

type Storage interface {
	GetBooksList() (list []models.Book, err error)
	DeleteBook(ID uint64) (err error)
	AddBook(book models.Book) (ID int, err error)
	UpdateBook(book models.Book) (err error)
}

type Service struct {
	log     zerolog.Logger
	storage Storage
}

func (s *Service) GetBooksList() (list []models.Book, err error) {
	return s.storage.GetBooksList()
}

func (s *Service) DeleteBook(ID uint64) (err error) {
	return s.storage.DeleteBook(ID)
}

func (s *Service) AddBook(book models.Book) (ID int, err error) {
	return s.storage.AddBook(book)
}

func (s *Service) UpdateBook(book models.Book) (err error) {
	return s.storage.UpdateBook(book)
}

func New(log zerolog.Logger, storage Storage) *Service {
	return &Service{
		log:     log,
		storage: storage,
	}
}
