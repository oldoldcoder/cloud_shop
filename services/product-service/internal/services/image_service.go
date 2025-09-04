package services

import (
	"oldoldcoder.com/cloudshop/product-service/internal/database"
	"oldoldcoder.com/cloudshop/product-service/internal/models"

	"gorm.io/gorm"
)

type ImageService struct{ db *gorm.DB }

func NewImageService() *ImageService { return &ImageService{db: database.GetDB()} }

func (s *ImageService) Create(img *models.ProductImage) error { return s.db.Create(img).Error }

func (s *ImageService) ListByProduct(productID uint) ([]models.ProductImage, error) {
	var list []models.ProductImage
	if err := s.db.Where("product_id = ?", productID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *ImageService) Delete(id uint) error { return s.db.Delete(&models.ProductImage{}, id).Error }
