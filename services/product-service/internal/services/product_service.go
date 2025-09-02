package services

import (
	"oldoldcoder.com/cloudshop/product-service/internal/database"
	"oldoldcoder.com/cloudshop/product-service/internal/models"

	"gorm.io/gorm"
)

// ProductService 商品服务
type ProductService struct {
	db *gorm.DB
}

// NewProductService 创建商品服务实例
func NewProductService() *ProductService {
	return &ProductService{
		db: database.GetDB(),
	}
}

// CreateProduct 创建商品
func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.db.Create(product).Error
}

// GetProductByID 根据ID获取商品
func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := s.db.Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProducts 获取商品列表
func (s *ProductService) GetProducts(page, pageSize int, categoryID *uint) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	query := s.db.Model(&models.Product{}).Preload("Category")
	
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// UpdateProduct 更新商品
func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.db.Save(product).Error
}

// DeleteProduct 删除商品
func (s *ProductService) DeleteProduct(id uint) error {
	return s.db.Delete(&models.Product{}, id).Error
}

// UpdateStock 更新库存
func (s *ProductService) UpdateStock(id uint, stock int) error {
	return s.db.Model(&models.Product{}).Where("id = ?", id).Update("stock", stock).Error
}
