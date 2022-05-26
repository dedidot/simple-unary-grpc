package repository

import "github.com/dedidot/grpc-book/model"

var storage []model.Book = []model.Book{}

func AddBook(bookData model.Book) model.Book {
	storage = append(storage, bookData)
	return bookData
}

func GetBook(bookId string) (int, model.Book) {
	for index, value := range storage {
		if value.Id == bookId {
			return index, value
		}
	}

	return 0, model.Book{}
}

func UpdateBook(bookData model.Book, bookId string) model.Book {
	index, book := GetBook(bookId)
	book.Id = bookData.Id
	book.Title = bookData.Title
	book.Author = bookData.Author
	book.IsRead = bookData.IsRead

	storage[index] = book
	return book
}

func DeleteBook(id string) bool {
	var afterDeleted []model.Book = []model.Book{}
	for _, book := range storage {
		if id != book.Id {
			afterDeleted = append(afterDeleted, book)
		}
	}

	storage = afterDeleted
	return true
}
