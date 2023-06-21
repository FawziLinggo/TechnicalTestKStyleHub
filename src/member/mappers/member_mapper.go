package mappers

import (
	"TechnicalTestKStyleHub/src/member"
)

type MemberMapper struct {
}

func NewMemberMapper() member.IMemberMapper {
	return &MemberMapper{}
}
