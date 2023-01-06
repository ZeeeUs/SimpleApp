package handlers

import (
	"github.com/ZeeeUs/SimpleApp/internal/models"
	"github.com/rs/zerolog"
)

type Book interface {
	GetBooksList() (list []models.Book, err error)
}

type Handler struct {
	log         zerolog.Logger
	BookService Book
}

func (h *Handler) WithBook(service Book) *Handler {
	h.BookService = service
	return h
}

func New(log zerolog.Logger) *Handler {
	return &Handler{
		log: log,
	}
}
