package models

type Book struct {
	ID        uint64 `json:"ID"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	ISBN      string `json:"ISBN"`
}
