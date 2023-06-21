package routes

import (
	"TechnicalTestKStyleHub/src/product/delivery/http/handlers"
	"github.com/labstack/echo/v4"
)

type ProductRoute struct {
	app            *echo.Echo
	ProductHandler handlers.ProductHandler
}

func NewProductRoute(app *echo.Echo, ProductHandler handlers.ProductHandler) ProductRoute {
	return ProductRoute{
		app:            app,
		ProductHandler: ProductHandler,
	}
}

func (route ProductRoute) RegisterRoutes() {

	// root
	productRoute := route.app.Group("/")

	// route
	productRoute.GET("product", route.ProductHandler.GetProduct)
	productRoute.GET("product/:id", route.ProductHandler.GetProductByID)
	productRoute.POST("product", route.ProductHandler.CreateProduct)
	productRoute.PUT("product", route.ProductHandler.EditProductByID)
	productRoute.DELETE("product", route.ProductHandler.DeleteProductByID)
	productRoute.GET("product/review", route.ProductHandler.GetProductReview)
	productRoute.POST("product/review", route.ProductHandler.CreateProductReview)
	productRoute.PUT("product/review", route.ProductHandler.EditProductReviewByID)
	productRoute.DELETE("product/review", route.ProductHandler.DeleteProductReviewByID)
	productRoute.POST("product/review/like", route.ProductHandler.CreateLikeReview)
	productRoute.DELETE("product/review/cancel-like", route.ProductHandler.CancelLikeReview)
}
