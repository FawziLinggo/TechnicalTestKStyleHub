package mappers

import (
	"TechnicalTestKStyleHub/domain/models"
	"TechnicalTestKStyleHub/domain/presenters"
	"TechnicalTestKStyleHub/src/product"
)

type ProductMapper struct {
}

func NewProductMapper() product.IProductMapper {
	return &ProductMapper{}
}

func (mapper ProductMapper) ToGetProductByIDResponse(data models.Product) (res presenters.GetProductByIDResponse) {
	var username []string
	var gender []string
	var skinType []string
	var skinColor []string
	var descReview []string
	jumlahReview := 0

	for _, v := range data.ProductReview {
		username = append(username, v.Member.Username)
		gender = append(gender, v.Member.Gender)
		skinType = append(skinType, v.Member.Skintype)
		skinColor = append(skinColor, v.Member.Skincolor)
		descReview = append(descReview, v.DescReview)
		jumlahReview = jumlahReview + len(v.LikeReview)
	}

	res = presenters.GetProductByIDResponse{
		ID:               data.ID,
		Username:         username,
		Gender:           gender,
		SkinType:         skinType,
		SkinColor:        skinColor,
		DescReview:       descReview,
		JumlahLikeReview: jumlahReview,
	}
	return res
}
