package entities

import "day2-task1/models"

type Books struct {
	Id        int    `json:"id_book"`
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required"`
	Publisher string `json:"publisher" validate:"required"`
}

type BooksDTO struct {
	Id        int
	Title     string
	Author    string
	Publisher string
}

//assembly DTO to model
func AssembBooksDTO(b *BooksDTO) *models.Books {
	return &models.Books{
		Title:     b.Title,
		Author:    b.Author,
		Publisher: b.Publisher,
	}
}

func ToBooksDTO(book *Books) *BooksDTO {
	return &BooksDTO{
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
	}
}
