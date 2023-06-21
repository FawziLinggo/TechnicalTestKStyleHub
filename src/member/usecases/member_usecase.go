package usecases

import (
	"TechnicalTestKStyleHub/app/constants"
	"TechnicalTestKStyleHub/app/helpers"
	"TechnicalTestKStyleHub/domain/models"
	"TechnicalTestKStyleHub/domain/presenters"
	"TechnicalTestKStyleHub/src/member"
	"net/http"
	"time"
)

type MemberUsecase struct {
	iMemberMapper     member.IMemberMapper
	iMemberRepository member.IMemberRepository
}

func NewMemberUseCase(
	iMemberMapper member.IMemberMapper,
	iMemberRepository member.IMemberRepository,
) member.IMemberUseCase {
	return &MemberUsecase{
		iMemberMapper:     iMemberMapper,
		iMemberRepository: iMemberRepository,
	}
}

func (uc MemberUsecase) GetMember() (res []presenters.GetMemberResponse, errDetail presenters.ErrorDetail) {

	// repo
	res, err := uc.iMemberRepository.GetMembers()
	if err != nil {
		return res, helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	return res, helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc MemberUsecase) CreateMember(input *presenters.CreateMemberRequest) (errDetail presenters.ErrorDetail) {

	//  model
	now := time.Now()
	data := models.Member{
		Username:  input.Username,
		Gender:    input.Gender,
		Skincolor: input.SkinColor,
		Skintype:  input.SkinType,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// repo
	err := uc.iMemberRepository.CreateMember(&data)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}

	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}

func (uc MemberUsecase) EditMemberByID(input *presenters.EditMemberByIDRequest) (errDetail presenters.ErrorDetail) {
	// model
	now := time.Now()

	data := models.Member{
		Username:  input.Username,
		Gender:    input.Gender,
		Skincolor: input.SkinColor,
		Skintype:  input.SkinType,
		UpdatedAt: now,
	}
	err := uc.iMemberRepository.EditMemberByID(input.ID, &data)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}
	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)

}

func (uc MemberUsecase) DeleteMemberByID(id string) (errDetail presenters.ErrorDetail) {
	err := uc.iMemberRepository.DeleteMemberByID(id)
	if err != nil {
		return helpers.NewBasicErrorDetail(http.StatusInternalServerError, err.Error(), err)
	}
	return helpers.NewBasicErrorDetail(http.StatusOK, constants.MessageSuccess, nil)
}
