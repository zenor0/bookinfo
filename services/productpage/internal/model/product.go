package model

// BookDetails 包含从 details 服务获取的图书详情
type BookDetails struct {
	ID          uint    `json:"id"`
	ISBN        string  `json:"isbn"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Publisher   string  `json:"publisher"`
	Language    string  `json:"language"`
	Pages       int     `json:"pages"`
	Year        int     `json:"year"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CoverImage  string  `json:"cover_image"`
}

// Review 包含从 reviews 服务获取的评论信息
type Review struct {
	ID         uint    `json:"id"`
	BookID     uint    `json:"book_id"`
	UserID     uint    `json:"user_id"`
	Username   string  `json:"username"`
	Content    string  `json:"content"`
	Rating     float64 `json:"rating,omitempty"`
	RatingHTML string  `json:"rating_html,omitempty"`
	Color      string  `json:"color,omitempty"`
}

// Rating 包含从 ratings 服务获取的评分信息
type Rating struct {
	BookID        uint    `json:"book_id"`
	AverageRating float64 `json:"average_rating"`
}

// ProductPage 包含完整的产品页面信息
type ProductPage struct {
	Book    *BookDetails `json:"book"`
	Reviews []Review     `json:"reviews"`
	Rating  *Rating      `json:"rating,omitempty"`
	User    string       `json:"user,omitempty"`
	Error   string       `json:"error,omitempty"`
	Version string       `json:"version,omitempty"`
}
