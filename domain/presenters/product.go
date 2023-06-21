package presenters

type GetProductResponse struct {
	ID    string `json:"id" example:"c836c745-e1db-4777-9699-e40fffaa302d"`
	Name  string `json:"name" example:"Product 1"`
	Price string `json:"price" example:"100000"`
}

type GetProductByIDResponse struct {
	ID               string       `json:"id" example:"c836c745-e1db-4777-9699-e40fffaa302d"`
	DataReviews      []DataReview `json:"data_reviews"`
	TotalLikeReviews int          `json:"total_like_review" example:"100"`
}

type GetProductReviewResponse struct {
	ID         string `json:"id" example:"c836c745-e1db-4777-9699-e40fffaa302d"`
	ProductID  string `json:"product_id" example:"c836c745-e1db-4777-9699-e40fffaa302d"`
	MemberID   string `json:"member_id" example:"c836c745-e1db-4777-9699-e40fffaa302d"`
	DescReview string `json:"desc_review" example:"Product 1"`
}

type DataReview struct {
	Username   string `json:"username" example:"fawzi"`
	Gender     string `json:"gender" example:"male"`
	SkinType   string `json:"skin_type" example:"oily"`
	SkinColor  string `json:"skin_color" example:"light"`
	DescReview string `json:"desc_review" example:"It's good"`
	TotalLike  int    `json:"total_like" example:"1"`
}
