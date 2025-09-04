package services

import (
	"oldoldcoder.com/cloudshop/product-service/internal/database"
	"oldoldcoder.com/cloudshop/product-service/internal/models"

	"gorm.io/gorm"
)

type SKUService struct{ db *gorm.DB }

func NewSKUService() *SKUService { return &SKUService{db: database.GetDB()} }

func (s *SKUService) Create(sku *models.ProductSKU) error { return s.db.Create(sku).Error }

func (s *SKUService) GetByID(id uint) (*models.ProductSKU, error) {
	var m models.ProductSKU
	if err := s.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *SKUService) ListByProduct(productID uint) ([]models.ProductSKU, error) {
	var list []models.ProductSKU
	if err := s.db.Where("product_id = ?", productID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *SKUService) Update(sku *models.ProductSKU) error { return s.db.Save(sku).Error }

func (s *SKUService) Delete(id uint) error { return s.db.Delete(&models.ProductSKU{}, id).Error }
