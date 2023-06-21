package member

import (
	"TechnicalTestKStyleHub/domain/models"
	"TechnicalTestKStyleHub/domain/presenters"
)

type IMemberRepository interface {
	GetMembers() (data []presenters.GetMemberResponse, err error)
	GetMemberByID(MemberID string) (data models.Member, err error)
	CreateMember(data *models.Member) (err error)
	EditMemberByID(MemberID string, data *models.Member) (err error)
	DeleteMemberByID(MemberID string) (err error)
}
