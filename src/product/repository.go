package product

import (
	"TechnicalTestKStyleHub/domain/models"
	"TechnicalTestKStyleHub/domain/presenters"
)

type IProductRepository interface {
	GetProducts() (data []presenters.GetProductResponse, err error)
	GetProductByID(ProductID string) (data models.Product, err error)
	CreateProduct(data *models.Product) (err error)
	EditProductByID(ProductID string, data *models.Product) (err error)
	DeleteProductByID(ProductID string) (err error)
	GetProductReviews() (data []presenters.GetProductReviewResponse, err error)
	GetProductReviewByID(ProductReviewID string) (data models.ProductReview, err error)
	CreateProductReview(data *models.ProductReview) (err error)
	EditProductReviewByID(ProductReviewID string, data *models.ProductReview) (err error)
	DeleteProductReviewByID(ProductReviewID string) (err error)
	CreateLikeReview(data *models.LikeReview) (err error)
	CancelLikeReview(likeReviewID string) (err error)
}
