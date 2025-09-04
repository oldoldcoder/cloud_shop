package services

import (
	"oldoldcoder.com/cloudshop/product-service/internal/database"
	"oldoldcoder.com/cloudshop/product-service/internal/models"

	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{db: database.GetDB()}
}

func (s *CategoryService) Create(category *models.Category) error {
	return s.db.Create(category).Error
}

func (s *CategoryService) GetByID(id uint) (*models.Category, error) {
	var m models.Category
	if err := s.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *CategoryService) List() ([]models.Category, error) {
	var list []models.Category
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *CategoryService) Update(category *models.Category) error {
	return s.db.Save(category).Error
}

func (s *CategoryService) Delete(id uint) error {
	return s.db.Delete(&models.Category{}, id).Error
}
