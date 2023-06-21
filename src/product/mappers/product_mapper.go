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
	var descReview presenters.DataReview
	var descReviews []presenters.DataReview
	jumlahReview := 0

	for _, v := range data.ProductReview {
		descReview = presenters.DataReview{
			Username:   v.Member.Username,
			Gender:     v.Member.Gender,
			SkinType:   v.Member.Skintype,
			SkinColor:  v.Member.Skincolor,
			DescReview: v.DescReview,
			TotalLike:  len(v.LikeReview),
		}
		descReviews = append(descReviews, descReview)
		jumlahReview = jumlahReview + len(v.LikeReview)

	}

	res = presenters.GetProductByIDResponse{
		ID:               data.ID,
		Name:             data.Name,
		Price:            data.Price,
		DataReviews:      descReviews,
		TotalLikeReviews: jumlahReview,
	}
	return res
}
