package presenters

import "github.com/go-playground/validator/v10"

// BasicResponse
// @Description General basic response
type BasicResponse struct {
	Status  bool               `json:"status" example:"false"`                  // Status turn true if process SUCCESS
	Error   string             `json:"error,omitempty" example:"ERROR or NULL"` // Error will be empty if SUCCESS
	Message string             `json:"message" example:"FAILED or SUCCESS and etc"`
	Data    CreateEditResponse `json:"data"`
}

// PaginationRequest
// @Description General basic pagination request
type PaginationRequest struct {
	Search string `query:"search"` // Search query
	Limit  int    `query:"limit"`  // Limit page limit
	Page   int    `query:"page"`   // Page number
}

// PaginationResponse
// @Description General basic pagination response, only data is dynamic array of object
type PaginationResponse struct {
	Total       int         `json:"total" example:"50"` // Total is total number of data
	CurrentPage int         `json:"current_page" example:"2"`
	PerPage     int         `json:"per_page" example:"5"`
	LastPage    int         `json:"last_page" example:"10"`
	From        int         `json:"from" example:"1"` // From is data first index
	To          int         `json:"to" example:"5"`   // To is data last index
	Data        interface{} `json:"data"`             // Data is dynamic array of object
	Other       interface{} `json:"other"`
}

type CreateEditResponse struct {
	ID    string      `json:"id"`
	Value interface{} `json:"value"`
}

type ErrorDetail struct {
	Code    int
	Error   error
	Message string
	Data    CreateEditResponse
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
