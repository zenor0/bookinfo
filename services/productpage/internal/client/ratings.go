package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zenor0/bookinfo/services/productpage/internal/model"
)

type RatingsClient struct {
	baseURL    string
	httpClient *http.Client
}

type RatingResponse struct {
	AverageRating float64 `json:"average_rating"`
}

func NewRatingsClient(baseURL string) *RatingsClient {
	return &RatingsClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *RatingsClient) GetBookRating(bookID uint) (*model.Rating, error) {
	url := fmt.Sprintf("%s/api/v1/products/%d/average-rating", c.baseURL, bookID)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get rating: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ratings service returned status code: %d", resp.StatusCode)
	}

	var ratingResp RatingResponse
	if err := json.NewDecoder(resp.Body).Decode(&ratingResp); err != nil {
		return nil, fmt.Errorf("failed to decode rating: %w", err)
	}

	return &model.Rating{
		BookID:        bookID,
		AverageRating: ratingResp.AverageRating,
	}, nil
}
