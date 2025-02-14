package client

import (
	"context"
	"fmt"

	pb "github.com/zenor0/bookinfo/proto/reviews"
	"github.com/zenor0/bookinfo/services/productpage/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type ReviewsClient struct {
	client pb.ReviewServiceClient
	conn   *grpc.ClientConn
}

func NewReviewsClient(target string) (*ReviewsClient, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to reviews service: %w", err)
	}

	return &ReviewsClient{
		client: pb.NewReviewServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *ReviewsClient) Close() error {
	return c.conn.Close()
}

func (c *ReviewsClient) GetBookReviews(bookID uint, user string) ([]model.Review, error) {
	ctx := context.Background()
	if user != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, "end-user", user)
	}

	resp, err := c.client.GetBookReviews(ctx, &pb.GetBookReviewsRequest{
		BookId:   uint32(bookID),
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get reviews: %w", err)
	}

	var reviews []model.Review
	for _, review := range resp.Data {
		reviews = append(reviews, model.Review{
			ID:         uint(review.Review.Id),
			BookID:     uint(review.Review.BookId),
			UserID:     uint(review.Review.UserId),
			Username:   review.Review.Username,
			Content:    review.Review.Content,
			Rating:     review.Rating,
			RatingHTML: review.RatingHtml,
			Color:      review.Color,
		})
	}

	return reviews, nil
}
