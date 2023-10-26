package postgresql

import (
	"advanced-rest-yt/internal/author/model"
	"database/sql"
	"time"
)

type Author struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Age       sql.NullInt16 `json:"age"`
	CreatedAt time.Time     `json:"created_at"`
	IsAlive   sql.NullBool  `json:"is_alive"`
}

func (a *Author) ToDomain() (m model.Author) {
	m = model.Author{
		ID:   a.ID,
		Name: a.Name,
	}

	if a.Age.Valid {
		m.Age = int(a.Age.Int16)
	}

	if a.IsAlive.Valid {
		m.IsAlive = a.IsAlive.Bool
	}

	return m
}
