package repositories

import (
	"TechnicalTestKStyleHub/domain/models"
	"TechnicalTestKStyleHub/domain/presenters"
	"TechnicalTestKStyleHub/src/product"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ProductRepository struct {
	sqlDB *gorm.DB
}

func NewProductRepository(sqlDB *gorm.DB) product.IProductRepository {
	return &ProductRepository{sqlDB: sqlDB}
}

func (repo ProductRepository) GetProducts() (data []presenters.GetProductResponse, err error) {
	var data2 []models.Product

	// using gorm
	if res := repo.sqlDB.Model(models.Product{}).Find(&data2); res.Error != nil {
		return data, res.Error
	}
	for _, v := range data2 {
		data = append(data, presenters.GetProductResponse{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
		})
	}

	return data, err
}

func (repo ProductRepository) GetProductByID(ProductID string) (data models.Product, err error) {
	query := repo.sqlDB.Preload("ProductReview").Preload("ProductReview.LikeReview").
		Preload("ProductReview.Member").
		Where("id = ?", ProductID).First(&data)

	if query.Error != nil {
		return data, err
	}

	return data, err
}

func (repo ProductRepository) CreateProduct(data *models.Product) (err error) {
	if res := repo.sqlDB.Create(&data); res.Error != nil {
		return res.Error
	}
	return err
}

func (repo ProductRepository) EditProductByID(ProductID string, data *models.Product) (err error) {
	// check if data is not exist
	queryCheck := repo.sqlDB.Where("id = ?", ProductID).First(&models.Product{})

	if queryCheck.Error != nil {
		return queryCheck.Error
	}

	if res := repo.sqlDB.Where("id = ?", ProductID).Updates(&data); res.Error != nil {
		return res.Error
	}
	return err
}

func (repo ProductRepository) DeleteProductByID(ProductID string) (err error) {
	queryCheck := repo.sqlDB.Where("id = ?", ProductID).First(&models.Product{})

	if queryCheck.Error != nil {
		return queryCheck.Error
	}

	if res := repo.sqlDB.Where("id = ?", ProductID).Delete(&models.Product{}); res.Error != nil {
		return res.Error
	}
	return err
}

func (repo ProductRepository) GetProductReviews() (data []presenters.GetProductReviewResponse, err error) {
	var data2 []models.ProductReview

	if res := repo.sqlDB.Model(models.ProductReview{}).Find(&data2); res.Error != nil {
		return data, res.Error
	}
	for _, v := range data2 {
		data = append(data, presenters.GetProductReviewResponse{
			ID:         v.ID,
			ProductID:  v.ProductID,
			MemberID:   v.MemberID,
			DescReview: v.DescReview,
		})
	}

	return data, err
}

func (repo ProductRepository) GetProductReviewByID(ProductReviewID string) (data models.ProductReview, err error) {
	if res := repo.sqlDB.Where("id = ?", ProductReviewID).First(&data); res.Error != nil {
		return data, res.Error
	}
	return data, err
}

func (repo ProductRepository) CreateProductReview(data *models.ProductReview) (err error) {
	if res := repo.sqlDB.Create(&data); res.Error != nil {
		return res.Error
	}
	return err
}

func (repo ProductRepository) EditProductReviewByID(ProductReviewID string, data *models.ProductReview) (err error) {
	queryCheck := repo.sqlDB.Where("id = ?", ProductReviewID).First(&models.ProductReview{})

	if queryCheck.Error != nil {
		return queryCheck.Error
	}

	if res := repo.sqlDB.Where("id = ?", ProductReviewID).Updates(&data); res.Error != nil {
		return res.Error
	}
	return err
}

func (repo ProductRepository) DeleteProductReviewByID(ProductReviewID string) error {
	var productReview models.ProductReview
	if err := repo.sqlDB.First(&productReview, "id = ?", ProductReviewID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("product review not found")
		}
		return err
	}

	if err := repo.sqlDB.Model(&productReview).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (repo ProductRepository) CreateLikeReview(data *models.LikeReview) (err error) {
	if res := repo.sqlDB.Create(&data); res.Error != nil {
		return res.Error
	}
	return err
}

func (repo ProductRepository) CancelLikeReview(likeReviewID string) error {
	var likeReview models.LikeReview
	if err := repo.sqlDB.First(&likeReview, "id = ?", likeReviewID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("like review not found")
		}
		return err
	}

	if err := repo.sqlDB.Model(&likeReview).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
