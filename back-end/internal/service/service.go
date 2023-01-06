package service

import (
	"github.com/ZeeeUs/SimpleApp/internal/models"
	"github.com/rs/zerolog"
)

type Storage interface {
	GetBooksList() (list []models.Book, err error)
}

type Service struct {
	log     zerolog.Logger
	storage Storage
}

func (s *Service) GetBooksList() (list []models.Book, err error) {
	return []models.Book{
		models.Book{
			ID:        1,
			Name:      "First Book",
			Author:    "Pushkin",
			Publisher: "ALPINA",
			ISBN:      "777-777-777",
		},
	}, nil
}

func New(log zerolog.Logger, storage Storage) *Service {
	return &Service{
		log:     log,
		storage: storage,
	}
}
