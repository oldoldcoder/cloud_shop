package services

import (
	"oldoldcoder.com/cloudshop/product-service/internal/database"
	"oldoldcoder.com/cloudshop/product-service/internal/models"

	"gorm.io/gorm"
)

type StockService struct{ db *gorm.DB }

func NewStockService() *StockService { return &StockService{db: database.GetDB()} }

func (s *StockService) GetBySKU(skuID uint) (*models.ProductStock, error) {
	var m models.ProductStock
	if err := s.db.Where("sku_id = ?", skuID).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *StockService) Upsert(stock *models.ProductStock) error { return s.db.Save(stock).Error }

func (s *StockService) Increase(skuID uint, delta int) error {
	return s.db.Model(&models.ProductStock{}).Where("sku_id = ?", skuID).
		UpdateColumn("quantity", gorm.Expr("quantity + ?", delta)).Error
}

func (s *StockService) Decrease(skuID uint, delta int) error {
	return s.db.Model(&models.ProductStock{}).Where("sku_id = ? AND quantity >= ?", skuID, delta).
		UpdateColumn("quantity", gorm.Expr("quantity - ?", delta)).Error
}
