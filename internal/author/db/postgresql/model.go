package postgresql

import (
	"advanced-rest-yt/internal/author/model"
	"database/sql"
)

type Author struct {
	ID   string        `json:"id"`
	Name string        `json:"name"`
	Age  sql.NullInt16 `json:"age"`
}

func (a *Author) ToDomain() (m model.Author) {
	m = model.Author{
		ID:   a.ID,
		Name: a.Name,
	}

	if a.Age.Valid {
		m.Age = int(a.Age.Int16)
	}

	return m
}
