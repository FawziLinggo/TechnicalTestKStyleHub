package presenters

type CreateProductReviewRequest struct {
	ProductID  string `json:"product_id" validate:"required"`
	MemberID   string `json:"member_id" validate:"required"`
	DescReview string `json:"desc_review" validate:"required"`
}

type EditProductReviewByIDRequest struct {
	ID         string `json:"id" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	MemberID   string `json:"member_id" validate:"required"`
	DescReview string `json:"desc_review" validate:"required"`
}

type DeleteProductReviewByIDRequest struct {
	ID string `json:"id" validate:"required"`
}
