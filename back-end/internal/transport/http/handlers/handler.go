package handlers

import (
	"github.com/ZeeeUs/SimpleApp/internal/models"
	"github.com/rs/zerolog"
)

type BookService interface {
	GetBooksList() (list []models.Book, err error)
	DeleteBook(ID uint64) (err error)
	AddBook(book models.Book) (ID int, err error)
}

type Handler struct {
	log         zerolog.Logger
	BookService BookService
}

func (h *Handler) WithBook(service BookService) *Handler {
	h.BookService = service
	return h
}

func New(log zerolog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}
