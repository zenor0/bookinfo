package server

import (
	"context"
	"strconv"

	pb "github.com/zenor0/bookinfo/proto/details"
	"github.com/zenor0/bookinfo/services/details/internal/model"
	"github.com/zenor0/bookinfo/services/details/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookServer struct {
	pb.UnimplementedBookServiceServer
	service *service.BookService
}

func NewBookServer(service *service.BookService) *BookServer {
	return &BookServer{service: service}
}

func (s *BookServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.Book, error) {
	year, _ := strconv.Atoi(req.Book.Year)
	book := &model.Book{
		Title:       req.Book.Title,
		Author:      req.Book.Author,
		ISBN:        req.Book.Isbn,
		Publisher:   req.Book.Publisher,
		Language:    req.Book.Language,
		Pages:       int(req.Book.Pages),
		Year:        year,
		Price:       float64(req.Book.Price),
		Description: req.Book.Description,
		CoverImage:  req.Book.CoverImage,
	}

	if err := s.service.CreateBook(book); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return convertToProtoBook(book), nil
}

func (s *BookServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.Book, error) {
	book, err := s.service.GetBook(uint(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "book not found")
	}

	return convertToProtoBook(book), nil
}

func (s *BookServer) GetBookByISBN(ctx context.Context, req *pb.GetBookByISBNRequest) (*pb.Book, error) {
	book, err := s.service.GetBookByISBN(req.Isbn)
	if err != nil {
		return nil, status.Error(codes.NotFound, "book not found")
	}

	return convertToProtoBook(book), nil
}

func (s *BookServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	books, total, err := s.service.ListBooks(int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var protoBooks []*pb.Book
	for _, book := range books {
		protoBooks = append(protoBooks, convertToProtoBook(&book))
	}

	return &pb.ListBooksResponse{
		Data: protoBooks,
		Meta: &pb.ListBooksResponse_Meta{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    int32(total),
		},
	}, nil
}

func (s *BookServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.Book, error) {
	year, _ := strconv.Atoi(req.Book.Year)
	book := &model.Book{
		ID:          uint(req.Book.Id),
		Title:       req.Book.Title,
		Author:      req.Book.Author,
		ISBN:        req.Book.Isbn,
		Publisher:   req.Book.Publisher,
		Language:    req.Book.Language,
		Pages:       int(req.Book.Pages),
		Year:        year,
		Price:       float64(req.Book.Price),
		Description: req.Book.Description,
		CoverImage:  req.Book.CoverImage,
	}

	if err := s.service.UpdateBook(book); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return convertToProtoBook(book), nil
}

func (s *BookServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.Empty, error) {
	if err := s.service.DeleteBook(uint(req.Id)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func (s *BookServer) SearchBooks(ctx context.Context, req *pb.SearchBooksRequest) (*pb.SearchBooksResponse, error) {
	books, err := s.service.SearchBooks(req.Query)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var protoBooks []*pb.Book
	for _, book := range books {
		protoBooks = append(protoBooks, convertToProtoBook(&book))
	}

	return &pb.SearchBooksResponse{
		Books: protoBooks,
	}, nil
}

func convertToProtoBook(book *model.Book) *pb.Book {
	return &pb.Book{
		Id:          uint32(book.ID),
		Title:       book.Title,
		Author:      book.Author,
		Isbn:        book.ISBN,
		Publisher:   book.Publisher,
		Language:    book.Language,
		Pages:       int32(book.Pages),
		Year:        strconv.Itoa(book.Year),
		Price:       float32(book.Price),
		Description: book.Description,
		CoverImage:  book.CoverImage,
		CreatedAt:   book.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   book.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
