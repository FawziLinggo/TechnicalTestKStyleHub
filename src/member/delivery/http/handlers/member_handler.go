package handlers

import (
	"TechnicalTestKStyleHub/app/constants"
	"TechnicalTestKStyleHub/app/helpers"
	"TechnicalTestKStyleHub/domain/presenters"
	"TechnicalTestKStyleHub/src/member"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MemberHandler struct {
	memberUseCase member.IMemberUseCase
}

func NewMemberHandler(IMemberUseCase member.IMemberUseCase) MemberHandler {
	return MemberHandler{memberUseCase: IMemberUseCase}
}

// GetMember
// @Summary Get All Member
// @Description Get All Member
// @Tags Member
// @Accept json
// @Success 200 {object} []presenters.GetMemberResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /member [GET]
func (handler MemberHandler) GetMember(ctx echo.Context) error {

	// get Member data
	//_, ok := ctx.Locals(constants.ContextMemberData).(models.Members)
	//if !ok {
	//	return helpers.NewBasicResponse(ctx, fiber.StatusInternalServerError, constants.False, errors.New(constants.MessageInterfaceConversionError), constants.MessageMemberDataFailed)
	//}

	// uc
	res, errDetail := handler.memberUseCase.GetMember()
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, errDetail.Message)
	}

	return ctx.JSON(200, res)
}

// CreateMember
// @Summary Create Member
// @Description Create Member
// @Tags Member
// @Accept json
// @Param request body presenters.CreateMemberRequest true "Request"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /member [POST]
func (handler MemberHandler) CreateMember(ctx echo.Context) error {

	//input
	input := new(presenters.CreateMemberRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.memberUseCase.CreateMember(input)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// EditMemberByID
// @Summary Edit Member By ID
// @Description Edit Member By ID
// @Tags Member
// @Accept json
// @Param id path string true "id"
// @Param request body presenters.EditMemberByIDRequest true "Request"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /member [PUT]
func (handler MemberHandler) EditMemberByID(ctx echo.Context) error {

	//input
	input := new(presenters.EditMemberByIDRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.memberUseCase.EditMemberByID(input)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}

// DeleteMemberByID
// @Summary Delete Member By ID
// @Description Delete Member By ID
// @Tags Member
// @Accept json
// @Param id path string true "id"
// @Success 200 {object} presenters.BasicResponse
// @Failure 401 {object} presenters.BasicResponse
// @Failure 500 {object} presenters.BasicResponse
// @Router /member [DELETE]
func (handler MemberHandler) DeleteMemberByID(ctx echo.Context) error {

	//input
	input := new(presenters.DeleteMemberByIDRequest)
	if err := ctx.Bind(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	// Validate
	if err := ctx.Validate(input); err != nil {
		return helpers.NewBasicResponse(ctx, http.StatusBadRequest, constants.False, err, constants.MessageInvalidRequest)
	}

	//uc
	errDetail := handler.memberUseCase.DeleteMemberByID(input.ID)
	if errDetail.Error != nil {
		return helpers.NewBasicResponse(ctx, errDetail.Code, constants.False, errDetail.Error, constants.MessageInvalidRequest)
	}

	return helpers.NewBasicResponse(ctx, http.StatusOK, constants.True, nil, constants.MessageSuccess)
}
