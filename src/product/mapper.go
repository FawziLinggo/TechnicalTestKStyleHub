package product

import (
	"TechnicalTestKStyleHub/domain/models"
	"TechnicalTestKStyleHub/domain/presenters"
)

type IProductMapper interface {
	ToGetProductByIDResponse(data models.Product) (res presenters.GetProductByIDResponse)
}
