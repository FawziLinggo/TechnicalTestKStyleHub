package member

import (
	"TechnicalTestKStyleHub/domain/presenters"
)

type IMemberUseCase interface {
	GetMember() (res []presenters.GetMemberResponse, errDetail presenters.ErrorDetail)
	CreateMember(req *presenters.CreateMemberRequest) (errDetail presenters.ErrorDetail)
	EditMemberByID(req *presenters.EditMemberByIDRequest) (errDetail presenters.ErrorDetail)
	DeleteMemberByID(id string) (errDetail presenters.ErrorDetail)
}
