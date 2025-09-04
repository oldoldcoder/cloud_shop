package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Product 商品表（SPU）
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null;size:200"`
	Description string    `json:"description" gorm:"type:text"`
	Brand       string    `json:"brand" gorm:"size:100"`
	Price       float64   `json:"price" gorm:"not null;type:decimal(10,2)"`
	Status      int       `json:"status" gorm:"default:1;type:tinyint"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// 软删除如果需要，可打开
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName 指定表名
func (Product) TableName() string {
	return "products"
}

// Category 分类表
type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null;size:100"`
	ParentID  *uint     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Category) TableName() string {
	return "categories"
}

// ProductCategory 多对多关联表（中间表）
type ProductCategory struct {
	ProductID  uint `json:"product_id" gorm:"primaryKey"`
	CategoryID uint `json:"category_id" gorm:"primaryKey"`
}

// TableName 指定表名
func (ProductCategory) TableName() string {
	return "product_category"
}

// ProductImage 商品图片表
type ProductImage struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"not null;size:500"`
	IsMain    int       `json:"is_main" gorm:"default:0;type:tinyint"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName 指定表名
func (ProductImage) TableName() string {
	return "product_images"
}

// ProductSKU SKU 表
type ProductSKU struct {
	ID         uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID  uint           `json:"product_id" gorm:"not null"`
	SKUCode    string         `json:"sku_code" gorm:"not null;size:100;uniqueIndex"`
	Attributes datatypes.JSON `json:"attributes" gorm:"type:json"`
	Price      float64        `json:"price" gorm:"not null;type:decimal(10,2)"`
	Stock      int            `json:"stock" gorm:"default:0"`
	Status     int            `json:"status" gorm:"default:1;type:tinyint"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

// TableName 指定表名
func (ProductSKU) TableName() string {
	return "product_skus"
}

// ProductStock 商品库存表
type ProductStock struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SkuID     uint      `json:"sku_id" gorm:"not null"`
	Warehouse string    `json:"warehouse" gorm:"size:100;default:'default'"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (ProductStock) TableName() string {
	return "product_stock"
}

// ProductDiscount 商品折扣表
type ProductDiscount struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID     uint      `json:"product_id" gorm:"not null"`
	DiscountType  string    `json:"discount_type" gorm:"not null;size:50"`
	DiscountValue float64   `json:"discount_value" gorm:"not null;type:decimal(10,2)"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	CreatedAt     time.Time `json:"created_at"`
}

// TableName 指定表名
func (ProductDiscount) TableName() string {
	return "product_discounts"
}
