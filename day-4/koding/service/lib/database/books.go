package database

import (
	"day2-task1/models"
	"errors"
)

func GetBooks() (interface{}, error) {

	modelUser := []models.Books{
		{
			Id:        1,
			Title:     "How To Be An Programmer",
			Author:    "James Wiel",
			Publisher: "Angola",
		},
		{
			Id:        2,
			Title:     "Learn Golang Basic",
			Author:    "Adam Qui",
			Publisher: "Angola",
		},
		{
			Id:        3,
			Title:     "Mastering Database",
			Author:    "Abe Parker",
			Publisher: "Angola",
		},
	}
	return modelUser, nil
}

func GetBookById(id int) (interface{}, error) {
	//get books from func GetBooks
	books, _ := GetBooks()

	var searchBook models.Books
	//convert interface to struct book
	allBooks := books.([]models.Books)

	//detect is book already
	alreadyBook := false

	//search book by Id
	for _, book := range allBooks {
		if book.Id == id {
			searchBook = book
			alreadyBook = true
		}
	}

	if !(alreadyBook) {
		return nil, errors.New("book not found")
	}

	return searchBook, nil
}

func CreateBook(book *models.Books) (interface{}, error) {
	//get books from func GetBooks
	books, _ := GetBooks()

	//convert interface to struct book
	allBooks := books.([]models.Books)

	checkValue := allBooks[0].Id

	//cari id yg paling besar
	for _, b := range allBooks {
		if b.Id > checkValue {
			checkValue = b.Id
		}
	}

	return models.Books{
		Id:        checkValue + 1,
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
	}, nil
}

func UpdateBookById(id int, book *models.Books) (interface{}, error) {
	//get books from func GetBooks
	books, _ := GetBooks()

	var searchBook models.Books
	//convert interface to struct book
	allBooks := books.([]models.Books)

	//detect is book already
	alreadyBook := false

	//search book by Id
	for _, book := range allBooks {
		if book.Id == id {
			alreadyBook = true
		}
	}

	if !(alreadyBook) {
		return nil, errors.New("book not found")
	} else {
		searchBook = models.Books{
			Id:        id,
			Title:     book.Title,
			Author:    book.Author,
			Publisher: book.Publisher,
		}
	}

	return searchBook, nil
}
