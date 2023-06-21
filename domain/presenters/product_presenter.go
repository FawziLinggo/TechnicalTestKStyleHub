package presenters

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type EditProductByIDRequest struct {
	ID    string  `json:"id" validate:"required"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type DeleteProductByIDRequest struct {
	ID string `json:"id" validate:"required"`
}
