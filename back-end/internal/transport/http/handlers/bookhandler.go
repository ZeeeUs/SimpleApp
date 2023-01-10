package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ZeeeUs/SimpleApp/internal/models"
	"github.com/gorilla/mux"
)

func (h *Handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	list, err := h.BookService.GetBooksList()
	if err != nil {
		h.log.Err(err).Send()
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonList, _ := json.Marshal(list)
	_, err = w.Write(jsonList)
	if err != nil {
		h.log.Err(err).Send()
	}
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	bookID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		h.log.Error().Msgf("can't get book id from url: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = h.BookService.DeleteBook(bookID); err != nil {
		h.log.Error().Msgf("can't delete book by id: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.log.Error().Msgf("can't decode book from request: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := h.BookService.AddBook(book)
	if err != nil {
		h.log.Error().Msgf("can't add the book: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonID, err := json.Marshal(id)
	if err != nil {
		h.log.Error().Msgf("can't add the book: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonID)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var book models.Book

	bookID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		h.log.Error().Msgf("can't get book id from url: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&book); err != nil {
		h.log.Error().Msgf("can't decode book from request: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DON'T DO THAT ON PROD SERVICE
	// USE DTO
	book.ID = bookID

	err = h.BookService.UpdateBook(book)
	if err != nil {
		h.log.Error().Msgf("can't update the book: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
