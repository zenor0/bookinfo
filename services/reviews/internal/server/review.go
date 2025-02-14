package server

import (
	"context"

	pb "github.com/zenor0/bookinfo/proto/reviews"
	"github.com/zenor0/bookinfo/services/reviews/internal/model"
	"github.com/zenor0/bookinfo/services/reviews/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReviewServer struct {
	pb.UnimplementedReviewServiceServer
	service *service.ReviewService
}

func NewReviewServer(service *service.ReviewService) *ReviewServer {
	return &ReviewServer{service: service}
}

func (s *ReviewServer) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.Review, error) {
	review := &model.Review{
		BookID:   uint(req.Review.BookId),
		UserID:   uint(req.Review.UserId),
		Username: req.Review.Username,
		Content:  req.Review.Content,
		Rating:   req.Review.Rating,
	}

	if err := s.service.CreateReview(review); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return convertToProtoReview(review), nil
}

func (s *ReviewServer) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.Review, error) {
	reviewResp, err := s.service.GetReview(uint(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "review not found")
	}

	return convertToProtoReview(&reviewResp.Review), nil
}

func (s *ReviewServer) GetBookReviews(ctx context.Context, req *pb.GetBookReviewsRequest) (*pb.GetBookReviewsResponse, error) {
	reviews, total, err := s.service.GetBookReviews(uint(req.BookId), int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var protoReviews []*pb.ReviewResponse
	for _, review := range reviews {
		protoReviews = append(protoReviews, &pb.ReviewResponse{
			Review:     convertToProtoReview(&review.Review),
			Rating:     review.Rating,
			RatingHtml: review.RatingHTML,
			Color:      review.Color,
		})
	}

	return &pb.GetBookReviewsResponse{
		Data: protoReviews,
		Meta: &pb.GetBookReviewsResponse_Meta{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    int32(total),
		},
	}, nil
}

func (s *ReviewServer) GetUserReviews(ctx context.Context, req *pb.GetUserReviewsRequest) (*pb.GetUserReviewsResponse, error) {
	reviews, err := s.service.GetUserReviews(uint(req.UserId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var protoReviews []*pb.ReviewResponse
	for _, review := range reviews {
		protoReviews = append(protoReviews, &pb.ReviewResponse{
			Review:     convertToProtoReview(&review.Review),
			Rating:     review.Rating,
			RatingHtml: review.RatingHTML,
			Color:      review.Color,
		})
	}

	return &pb.GetUserReviewsResponse{
		Reviews: protoReviews,
	}, nil
}

func (s *ReviewServer) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.Review, error) {
	review := &model.Review{
		ID:       uint(req.Review.Id),
		BookID:   uint(req.Review.BookId),
		UserID:   uint(req.Review.UserId),
		Username: req.Review.Username,
		Content:  req.Review.Content,
		Rating:   req.Review.Rating,
	}

	if err := s.service.UpdateReview(review); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return convertToProtoReview(review), nil
}

func (s *ReviewServer) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.Empty, error) {
	if err := s.service.DeleteReview(uint(req.Id)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func convertToProtoReview(review *model.Review) *pb.Review {
	return &pb.Review{
		Id:        uint32(review.ID),
		BookId:    uint32(review.BookID),
		UserId:    uint32(review.UserID),
		Username:  review.Username,
		Content:   review.Content,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: review.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
