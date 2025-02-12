package client

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (c *RatingsClient) GetProductRating(productID string) (float64, error) {
	url := fmt.Sprintf("%s/api/v1/products/%s/average-rating", c.baseURL, productID)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to get rating: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("ratings service returned status code: %d", resp.StatusCode)
	}

	var rating RatingResponse
	if err := json.NewDecoder(resp.Body).Decode(&rating); err != nil {
		return 0, fmt.Errorf("failed to decode rating response: %w", err)
	}

	return rating.AverageRating, nil
}

// 不同版本的评分展示格式
type RatingFormatter interface {
	FormatRating(rating float64) (string, string)
}

// V1 版本：不显示评分
type V1Formatter struct{}

func (f *V1Formatter) FormatRating(rating float64) (string, string) {
	return "", ""
}

// V2 版本：黑色星星
type V2Formatter struct{}

func (f *V2Formatter) FormatRating(rating float64) (string, string) {
	stars := ""
	for i := 0; i < int(rating); i++ {
		stars += "★"
	}
	return stars, "black"
}

// V3 版本：红色星星
type V3Formatter struct{}

func (f *V3Formatter) FormatRating(rating float64) (string, string) {
	stars := ""
	for i := 0; i < int(rating); i++ {
		stars += "★"
	}
	return stars, "red"
}
