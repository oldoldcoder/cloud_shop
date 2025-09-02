package models

import (
	"time"
	"gorm.io/gorm"
)

// Product 商品模型
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null;size:255"`
	Description string         `json:"description" gorm:"type:text"`
	Price       float64        `json:"price" gorm:"not null;type:decimal(10,2)"`
	Stock       int            `json:"stock" gorm:"not null;default:0"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
	Status      int            `json:"status" gorm:"not null;default:1"` // 1: 上架, 0: 下架
	ImageURL    string         `json:"image_url" gorm:"size:500"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Category 商品分类模型
type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null;size:100"`
	ParentID  *uint          `json:"parent_id" gorm:"default:null"`
	Level     int            `json:"level" gorm:"not null;default:1"`
	Sort      int            `json:"sort" gorm:"not null;default:0"`
	Status    int            `json:"status" gorm:"not null;default:1"` // 1: 启用, 0: 禁用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// ProductImage 商品图片模型
type ProductImage struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	URL       string         `json:"url" gorm:"not null;size:500"`
	Sort      int            `json:"sort" gorm:"not null;default:0"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// ProductSpec 商品规格模型
type ProductSpec struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Name      string         `json:"name" gorm:"not null;size:100"`
	Value     string         `json:"value" gorm:"not null;size:255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
