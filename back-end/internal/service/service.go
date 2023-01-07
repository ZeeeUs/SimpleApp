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
	return s.storage.GetBooksList()
}

func New(log zerolog.Logger, storage Storage) *Service {
	return &Service{
		log:     log,
		storage: storage,
	}
}
