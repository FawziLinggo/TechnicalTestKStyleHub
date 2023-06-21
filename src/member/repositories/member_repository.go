package repositories

import (
	"TechnicalTestKStyleHub/domain/models"
	"TechnicalTestKStyleHub/domain/presenters"
	"TechnicalTestKStyleHub/src/member"
	"gorm.io/gorm"
)

type MemberRepository struct {
	sqlDB *gorm.DB
}

func NewMemberRepository(sqlDB *gorm.DB) member.IMemberRepository {
	return &MemberRepository{sqlDB: sqlDB}
}

func (repo MemberRepository) GetMembers() (data []presenters.GetMemberResponse, err error) {
	var data2 []models.Member

	// using gorm
	if res := repo.sqlDB.Model(models.Member{}).Find(&data2); res.Error != nil {
		return data, res.Error
	}
	for _, v := range data2 {
		data = append(data, presenters.GetMemberResponse{
			ID:        v.ID,
			Username:  v.Username,
			SkinColor: v.Skincolor,
			SkinType:  v.Skintype,
			Gender:    v.Gender,
		})
	}

	return data, err
}

func (repo MemberRepository) GetMemberByID(MemberID string) (data models.Member, err error) {
	if res := repo.sqlDB.Where("id = ?", MemberID).First(&data); res.Error != nil {
		return data, res.Error
	}
	return data, err
}

func (repo MemberRepository) CreateMember(data *models.Member) (err error) {
	if res := repo.sqlDB.Create(&data); res.Error != nil {
		return res.Error
	}
	return err
}

func (repo MemberRepository) EditMemberByID(MemberID string, data *models.Member) (err error) {
	// check if data is not exist
	queryCheck := repo.sqlDB.Where("id = ?", MemberID).First(&models.Member{})

	if queryCheck.Error != nil {
		return queryCheck.Error
	}

	if res := repo.sqlDB.Where("id = ?", MemberID).Updates(&data); res.Error != nil {
		return res.Error
	}
	return err
}

func (repo MemberRepository) DeleteMemberByID(MemberID string) (err error) {
	// check if data is not exist
	queryCheck := repo.sqlDB.Where("id = ?", MemberID).First(&models.Member{})

	if queryCheck.Error != nil {
		return queryCheck.Error
	}

	if res := repo.sqlDB.Where("id = ?", MemberID).Delete(&models.Member{}); res.Error != nil {
		return res.Error
	}
	return err
}
