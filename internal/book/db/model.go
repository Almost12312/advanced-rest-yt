package book

import (
	"advanced-rest-yt/internal/author"
	"advanced-rest-yt/internal/book"
	"database/sql"
)

type Book struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Age     sql.NullInt16   `json:"age"`
	Authors []author.Author `json:"author"`
}

func (m *Book) ToDomain() (b book.Book) {
	b = book.Book{
		ID:      m.ID,
		Name:    m.Name,
		Authors: m.Authors,
	}

	if m.Age.Valid {
		b.Age = int(m.Age.Int16)
	}

	return
}
