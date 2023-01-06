package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) BookHandler(w http.ResponseWriter, r *http.Request) {
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
