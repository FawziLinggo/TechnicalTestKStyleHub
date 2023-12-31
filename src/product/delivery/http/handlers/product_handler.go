package handlers

import (
	"TechnicalTestKStyleHub/app/constants"
	"TechnicalTestKStyleHub/app/helpers"
	"TechnicalTestKStyleHub/domain/presenters"
	"TechnicalTestKStyleHub/src/product"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductHandler struct {
	productUseCase product.IProductUseCase
}

func NewProductHandler(IProductUseCase product.IProductUseCase) ProductHandler {
	return ProductHandler{productUseCase: IProductUseCase}
}

// GetProduct
// @Summary Get All Product
// @Description Get All Product
// @Tags Product
// @Accept json
// @Success 200 {object} []presenters.GetProductResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product [GET]
func (handler ProductHandler) GetProduct(ctx echo.Context) error {
	// uc
	res, errDetail := handler.productUseCase.GetProduct()
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, errDetail.Message)
	}

	return ctx.JSON(200, res)
}

// GetProductByID
// @Summary Get Product By ID
// @Description Get Product By ID
// @Tags Product
// @Accept json
// @Param id path string true "id"
// @Success 200 {object} presenters.GetProductByIDResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product/{id} [GET]
func (handler ProductHandler) GetProductByID(ctx echo.Context) error {
	//input
	id := ctx.Param("id")

	//uc
	res, errDetail := handler.productUseCase.GetProductByID(id)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, errDetail.Message)
	}

	return ctx.JSON(200, res)
}

// CreateProduct
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Param request body presenters.CreateProductRequest true "Request"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product [POST]
func (handler ProductHandler) CreateProduct(ctx echo.Context) error {

	//input
	input := new(presenters.CreateProductRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.productUseCase.CreateProduct(input)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// EditProductByID
// @Summary Edit Product By ID
// @Description Edit Product By ID
// @Tags Product
// @Accept json
// @Param id path string true "id"
// @Param request body presenters.EditProductByIDRequest true "Request"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product [PUT]
func (handler ProductHandler) EditProductByID(ctx echo.Context) error {

	//input
	input := new(presenters.EditProductByIDRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.productUseCase.EditProductByID(input)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// DeleteProductByID
// @Summary Delete Product By ID
// @Description Delete Product By ID
// @Tags Product
// @Accept json
// @Param id path string true "id"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product [DELETE]
func (handler ProductHandler) DeleteProductByID(ctx echo.Context) error {

	//input
	input := new(presenters.DeleteProductByIDRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.productUseCase.DeleteProductByID(input.ID)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// GetProductReview
// @Summary Get All ProductReview
// @Description Get All ProductReview
// @Tags ProductReview
// @Accept json
// @Success 200 {object} []presenters.GetProductReviewResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product/review [GET]
func (handler ProductHandler) GetProductReview(ctx echo.Context) error {
	// uc
	res, errDetail := handler.productUseCase.GetProductReview()
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, errDetail.Message)
	}

	return ctx.JSON(200, res)
}

// CreateProductReview
// @Summary Create Product Review
// @Description Create Product Review
// @Tags ProductReview
// @Accept json
// @Param request body presenters.CreateProductReviewRequest true "Request"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product/reveiw [POST]
func (handler ProductHandler) CreateProductReview(ctx echo.Context) error {

	//input
	input := new(presenters.CreateProductReviewRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.productUseCase.CreateProductReview(input)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// EditProductReviewByID
// @Summary Edit ProductReview By ID
// @Description Edit ProductReview By ID
// @Tags ProductReview
// @Accept json
// @Param id path string true "id"
// @Param request body presenters.EditProductReviewByIDRequest true "Request"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product/review [PUT]
func (handler ProductHandler) EditProductReviewByID(ctx echo.Context) error {

	//input
	input := new(presenters.EditProductReviewByIDRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.productUseCase.EditProductReviewByID(input)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// DeleteProductReviewByID
// @Summary Delete ProductReview By ID
// @Description Delete ProductReview By ID
// @Tags ProductReview
// @Accept json
// @Param id path string true "id"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product/review [DELETE]
func (handler ProductHandler) DeleteProductReviewByID(ctx echo.Context) error {

	//input
	input := new(presenters.DeleteProductReviewByIDRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.productUseCase.DeleteProductReviewByID(input.ID)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// CreateLikeReview
// @Summary Create LikeReview
// @Description Create LikeReview
// @Tags LikeReview
// @Accept json
// @Param request body presenters.CreateLikeReviewRequest true "Request"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product/review/like [POST]
func (handler ProductHandler) CreateLikeReview(ctx echo.Context) error {

	//input
	input := new(presenters.CreateLikeReviewRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.productUseCase.CreateLikeReview(input)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// CancelLikeReview
// @Summary Delete LikeReview
// @Description Delete LikeReview
// @Tags LikeReview
// @Accept json
// @Param request body presenters.CancelLikeReviewRequest true "Request"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /product/review/cancel-like [DELETE]
func (handler ProductHandler) CancelLikeReview(ctx echo.Context) error {

	//input
	input := new(presenters.CancelLikeReviewRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.productUseCase.CancelLikeReview(input.ID)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}
