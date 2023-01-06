package storage

import (
	"github.com/ZeeeUs/SimpleApp/internal/models"

	"github.com/jackc/pgx"
)

type Storage struct {
	conn *pgx.Conn
}

func (s *Storage) GetBooksList() (list []models.Book, err error) {
	query := "SELECT id, name, author, publisher, isbn FROM public.Books"

	rows, err := s.conn.Query(query)
	if err != nil {
		return nil, err
	}

	list = make([]models.Book, 0)
	for rows.Next() {
		var book models.Book
		if err = rows.Scan(&book.ID, &book.Name, &book.Author, &book.Publisher, &book.ISBN); err != nil {
			return nil, err
		}
		list = append(list, book)
	}
	return
}

func New(conn *pgx.Conn) *Storage {
	return &Storage{
		conn: conn,
	}
}
