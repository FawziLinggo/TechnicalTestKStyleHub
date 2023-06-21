package helpers

import (
	"TechnicalTestKStyleHub/domain/presenters"
	"github.com/labstack/echo/v4"
)

func NewBasicErrorDetail(code int, message string, err error) presenters.ErrorDetail {
	return presenters.ErrorDetail{
		Code:    code,
		Error:   err,
		Message: message,
	}
}

func NewBasicErrorDetailWithData(code int, message string, err error, data presenters.CreateEditResponse) presenters.ErrorDetail {
	return presenters.ErrorDetail{
		Code:    code,
		Error:   err,
		Message: message,
		Data:    data,
	}
}

func NewBasicResponse(ctx echo.Context, code int, status bool, err error, message string) error {

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return ctx.JSON(code, presenters.BasicResponse{
		Status:  status,
		Error:   errMsg,
		Message: message,
	})
}
