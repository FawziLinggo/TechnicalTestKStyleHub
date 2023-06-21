package presenters

type CreateLikeReviewRequest struct {
	ProductReviewID string `json:"product_review_id" validate:"required"`
	MemberID        string `json:"member_id" validate:"required"`
}

type CancelLikeReviewRequest struct {
	ID string `json:"id" validate:"required"`
}
