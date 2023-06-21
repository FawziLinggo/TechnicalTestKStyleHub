package usecases

import (
	"TechnicalTestKStyleHub/app/constants"
	"TechnicalTestKStyleHub/app/helpers"
	"TechnicalTestKStyleHub/domain/models"
	"TechnicalTestKStyleHub/domain/presenters"
	product "TechnicalTestKStyleHub/src/product"
	"log"
	"net/http"
	"time"
)

type ProductUsecase struct {
	iProductMapper     product.IProductMapper
	iProductRepository product.IProductRepository
}

func NewProductUseCase(
	iProductMapper product.IProductMapper,
	iProductRepository product.IProductRepository,
) product.IProductUseCase {
	return &ProductUsecase{
		iProductMapper:     iProductMapper,
		iProductRepository: iProductRepository,
	}
}

func (uc ProductUsecase) GetProduct() (res []presenters.GetProductResponse, errDetail presenters.ErrorDetail) {

	// repo
	res, err := uc.iProductRepository.GetProducts()
	if err != nil {
		return res, helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	return res, helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc ProductUsecase) GetProductByID(productID string) (res presenters.GetProductByIDResponse, errDetail presenters.ErrorDetail) {

	// repo
	data, err := uc.iProductRepository.GetProductByID(productID)
	if err != nil {
		return res, helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	// mapper
	res = uc.iProductMapper.ToGetProductByIDResponse(data)
	log.Println("data", data)

	return res, helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc ProductUsecase) CreateProduct(input *presenters.CreateProductRequest) (errDetail presenters.ErrorDetail) {

	//  model
	now := time.Now()
	data := models.Product{
		Name:      input.Name,
		Price:     input.Price,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// repo
	err := uc.iProductRepository.CreateProduct(&data)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc ProductUsecase) EditProductByID(input *presenters.EditProductByIDRequest) (errDetail presenters.ErrorDetail) {
	// model
	now := time.Now()

	data := models.Product{
		Name:      input.Name,
		Price:     input.Price,
		UpdatedAt: now,
	}
	err := uc.iProductRepository.EditProductByID(input.ID, &data)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}
	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)

}

func (uc ProductUsecase) DeleteProductByID(id string) (errDetail presenters.ErrorDetail) {
	err := uc.iProductRepository.DeleteProductByID(id)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}
	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc ProductUsecase) GetProductReview() (res []presenters.GetProductReviewResponse, errDetail presenters.ErrorDetail) {

	// repo
	res, err := uc.iProductRepository.GetProductReviews()
	if err != nil {
		return res, helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	return res, helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc ProductUsecase) CreateProductReview(input *presenters.CreateProductReviewRequest) (errDetail presenters.ErrorDetail) {

	//  model
	now := time.Now()
	data := models.ProductReview{
		ProductID:  input.ProductID,
		MemberID:   input.MemberID,
		DescReview: input.DescReview,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// repo
	err := uc.iProductRepository.CreateProductReview(&data)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc ProductUsecase) EditProductReviewByID(input *presenters.EditProductReviewByIDRequest) (errDetail presenters.ErrorDetail) {
	// model
	now := time.Now()

	data := models.ProductReview{
		ID:         input.ID,
		ProductID:  input.ProductID,
		MemberID:   input.MemberID,
		DescReview: input.DescReview,
		UpdatedAt:  now,
	}
	err := uc.iProductRepository.EditProductReviewByID(input.ID, &data)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}
	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)

}

func (uc ProductUsecase) DeleteProductReviewByID(id string) (errDetail presenters.ErrorDetail) {
	err := uc.iProductRepository.DeleteProductReviewByID(id)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}
	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc ProductUsecase) CreateLikeReview(input *presenters.CreateLikeReviewRequest) (errDetail presenters.ErrorDetail) {

	//  model
	now := time.Now()
	data := models.LikeReview{
		ProductReviewID: input.ProductReviewID,
		MemberID:        input.MemberID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	// repo
	err := uc.iProductRepository.CreateLikeReview(&data)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc ProductUsecase) CancelLikeReview(likeReviewID string) (errDetail presenters.ErrorDetail) {

	// repo
	err := uc.iProductRepository.CancelLikeReview(likeReviewID)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}
