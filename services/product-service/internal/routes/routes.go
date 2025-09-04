package routes

import (
	"oldoldcoder.com/cloudshop/product-service/internal/handlers"
	"oldoldcoder.com/cloudshop/product-service/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine, jwtSecret string) {
	// 创建处理器
	productHandler := handlers.NewProductHandler()
	categoryHandler := handlers.NewCategoryHandler()
	imageHandler := handlers.NewImageHandler()
	skuHandler := handlers.NewSKUHandler()
	stockHandler := handlers.NewStockHandler()
	discountHandler := handlers.NewDiscountHandler()

	// API 路由组（受保护）
	api := r.Group("/api")
	{
		// 商品相关路由（受 Auth 保护）
		products := api.Group("/products", middleware.Auth(jwtSecret))
		{
			products.POST("", productHandler.CreateProduct)       // 创建商品
			products.GET("", productHandler.GetProducts)          // 获取商品列表
			products.GET("/:id", productHandler.GetProduct)       // 获取商品详情
			products.PUT("/:id", productHandler.UpdateProduct)    // 更新商品
			products.DELETE("/:id", productHandler.DeleteProduct) // 删除商品

			// 图片
			products.POST("/:id/images", imageHandler.Create)
			products.GET("/:id/images", imageHandler.ListByProduct)
			products.DELETE("/images/:id", imageHandler.Delete)

			// SKU
			products.POST("/:id/skus", skuHandler.Create)
			products.GET("/:id/skus", skuHandler.ListByProduct)
			products.GET("/skus/:id", skuHandler.Get)
			products.PUT("/skus/:id", skuHandler.Update)
			products.DELETE("/skus/:id", skuHandler.Delete)

			// 库存
			products.GET("/stocks/:sku_id", stockHandler.Get)
			products.POST("/stocks", stockHandler.Upsert)
			products.POST("/stocks/:sku_id/increase", stockHandler.Increase)
			products.POST("/stocks/:sku_id/decrease", stockHandler.Decrease)

			// 折扣
			products.POST("/:id/discounts", discountHandler.Create)
			products.GET("/:id/discounts", discountHandler.ListByProduct)
			products.GET("/discounts/:id", discountHandler.Get)
			products.PUT("/discounts/:id", discountHandler.Update)
			products.DELETE("/discounts/:id", discountHandler.Delete)
		}

		// 分类（受 Auth 保护）
		categories := api.Group("/categories", middleware.Auth(jwtSecret))
		{
			categories.POST("", categoryHandler.Create)
			categories.GET("", categoryHandler.List)
			categories.GET("/:id", categoryHandler.Get)
			categories.PUT("/:id", categoryHandler.Update)
			categories.DELETE("/:id", categoryHandler.Delete)
		}
	}

	// Swagger 文档（公开）
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
