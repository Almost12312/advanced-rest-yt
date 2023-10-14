package book

type CreateBookDTO struct {
	Name     string `json:"name"`
	AuthorId string `json:"author_id"`
}
