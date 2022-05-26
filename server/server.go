package server

import (
	"context"
	"github.com/dedidot/grpc-book/model"
	"github.com/dedidot/grpc-book/proto/bookpb"
	"github.com/dedidot/grpc-book/service"
)

type Server struct {
	bookpb.BookServiceServer
}

func mapToBook(book model.Book) *bookpb.Book {
	return &bookpb.Book{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
		IsRead: book.IsRead,
	}
}

func (s *Server) GetBookData(ctx context.Context, req *bookpb.GetBookRequest) (*bookpb.GetBookResponse, error) {
	var bookId string = req.GetId()
	_, result := service.GetBook(bookId)

	if len(result.Id) == 0 || result.Id != bookId {
		return &bookpb.GetBookResponse{
			Status: false,
			Data:   nil,
		}, nil
	}

	var bookData *bookpb.Book = &bookpb.Book{
		Id:     result.Id,
		Title:  result.Title,
		Author: result.Author,
		IsRead: result.IsRead,
	}

	return &bookpb.GetBookResponse{
		Status: true,
		Data:   bookData,
	}, nil
}

func (s *Server) AddBookData(ctx context.Context, req *bookpb.AddBookRequest) (*bookpb.AddBookResponse, error) {
	var bookRequest *bookpb.Book = req.GetBook()

	var bookData model.Book = model.Book{
		Id:     bookRequest.GetId(),
		Title:  bookRequest.GetTitle(),
		Author: bookRequest.GetAuthor(),
		IsRead: bookRequest.GetIsRead(),
	}

	var insertedBook model.Book = service.AddBook(bookData)
	return &bookpb.AddBookResponse{
		Status: true,
		Data:   mapToBook(insertedBook),
	}, nil
}

func (s *Server) UpdateBookData(ctx context.Context, req *bookpb.UpdateBookRequest) (*bookpb.UpdateBookResponse, error) {
	var bookRequest *bookpb.Book = req.GetBook()
	var bookData model.Book = model.Book{
		Id:     bookRequest.GetId(),
		Title:  bookRequest.GetTitle(),
		Author: bookRequest.GetAuthor(),
		IsRead: bookRequest.GetIsRead(),
	}

	updateBook := service.UpdateBook(bookData, bookRequest.Id)
	return &bookpb.UpdateBookResponse{
		Status: true,
		Data:   mapToBook(updateBook),
	}, nil
}

func (s *Server) DeleteBookData(ctx context.Context, req *bookpb.DeleteBookRequest) (*bookpb.DeleteBookResponse, error) {
	var bookId string = req.GetId()
	result := service.DeleteBook(bookId)
	return &bookpb.DeleteBookResponse{
		Status: result,
	}, nil
}
