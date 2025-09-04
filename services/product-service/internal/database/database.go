package database

import (
	"fmt"
	"log"

	"oldoldcoder.com/cloudshop/product-service/internal/config"
	"oldoldcoder.com/cloudshop/product-service/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase(cfg *config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db

	// 自动迁移数据库表（可控）
	if cfg.Migrate.Auto {
		if err := autoMigrate(); err != nil {
			return fmt.Errorf("failed to auto migrate: %v", err)
		}
	}

	log.Println("Database initialized successfully")
	return nil
}

// autoMigrate 自动迁移数据库表
func autoMigrate() error {
	return DB.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.ProductCategory{},
		&models.ProductImage{},
		&models.ProductSKU{},
		&models.ProductStock{},
		&models.ProductDiscount{},
	)
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
