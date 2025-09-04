package services

import (
	"oldoldcoder.com/cloudshop/product-service/internal/database"
	"oldoldcoder.com/cloudshop/product-service/internal/models"

	"gorm.io/gorm"
)

type DiscountService struct{ db *gorm.DB }

func NewDiscountService() *DiscountService { return &DiscountService{db: database.GetDB()} }

func (s *DiscountService) Create(d *models.ProductDiscount) error { return s.db.Create(d).Error }

func (s *DiscountService) GetByID(id uint) (*models.ProductDiscount, error) {
	var m models.ProductDiscount
	if err := s.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *DiscountService) ListByProduct(productID uint) ([]models.ProductDiscount, error) {
	var list []models.ProductDiscount
	if err := s.db.Where("product_id = ?", productID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *DiscountService) Update(d *models.ProductDiscount) error { return s.db.Save(d).Error }

func (s *DiscountService) Delete(id uint) error {
	return s.db.Delete(&models.ProductDiscount{}, id).Error
}
