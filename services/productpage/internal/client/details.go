package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zenor0/bookinfo/pkg/tracing"
	pb "github.com/zenor0/bookinfo/proto/details"
	"github.com/zenor0/bookinfo/services/productpage/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DetailsClient struct {
	client pb.BookServiceClient
	conn   *grpc.ClientConn
}

func NewDetailsClient(target string) (*DetailsClient, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to details service: %w", err)
	}

	return &DetailsClient{
		client: pb.NewBookServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *DetailsClient) Close() error {
	return c.conn.Close()
}

func (c *DetailsClient) GetBookDetails(ctx context.Context, id uint) (*model.BookDetails, error) {
	ctx, span := tracing.StartSpan(ctx, "details.GetBookDetails")
	defer span.End()

	resp, err := c.client.GetBook(ctx, &pb.GetBookRequest{Id: uint32(id)})
	if err != nil {
		return nil, fmt.Errorf("failed to get book details: %w", err)
	}

	year, _ := strconv.Atoi(resp.Year)
	return &model.BookDetails{
		ID:          uint(resp.Id),
		Title:       resp.Title,
		Author:      resp.Author,
		ISBN:        resp.Isbn,
		Publisher:   resp.Publisher,
		Language:    resp.Language,
		Pages:       int(resp.Pages),
		Year:        year,
		Price:       float64(resp.Price),
		Description: resp.Description,
		CoverImage:  resp.CoverImage,
	}, nil
}

func (c *DetailsClient) GetBookList(ctx context.Context, page, pageSize int) (*pb.ListBooksResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "details.GetBookList")
	defer span.End()

	resp, err := c.client.ListBooks(ctx, &pb.ListBooksRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get book list: %w", err)
	}
	return resp, nil
}

func (c *DetailsClient) SearchBooks(ctx context.Context, query string) (*pb.SearchBooksResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "details.SearchBooks")
	defer span.End()

	resp, err := c.client.SearchBooks(ctx, &pb.SearchBooksRequest{
		Query: query,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to search books: %w", err)
	}
	return resp, nil
}
