package product

import (
	"TechnicalTestKStyleHub/domain/presenters"
)

type IProductUseCase interface {
	GetProduct() (res []presenters.GetProductResponse, errDetail presenters.ErrorDetail)
	GetProductByID(id string) (res presenters.GetProductByIDResponse, errDetail presenters.ErrorDetail)
	CreateProduct(req *presenters.CreateProductRequest) (errDetail presenters.ErrorDetail)
	EditProductByID(req *presenters.EditProductByIDRequest) (errDetail presenters.ErrorDetail)
	DeleteProductByID(id string) (errDetail presenters.ErrorDetail)
	GetProductReview() (res []presenters.GetProductReviewResponse, errDetail presenters.ErrorDetail)
	CreateProductReview(req *presenters.CreateProductReviewRequest) (errDetail presenters.ErrorDetail)
	EditProductReviewByID(req *presenters.EditProductReviewByIDRequest) (errDetail presenters.ErrorDetail)
	DeleteProductReviewByID(id string) (errDetail presenters.ErrorDetail)
	CreateLikeReview(req *presenters.CreateLikeReviewRequest) (errDetail presenters.ErrorDetail)
	CancelLikeReview(likeReviewID string) (errDetail presenters.ErrorDetail)
}
