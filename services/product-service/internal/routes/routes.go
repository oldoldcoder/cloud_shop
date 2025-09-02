package routes

import (
	"oldoldcoder.com/cloudshop/product-service/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine) {
	// 创建处理器
	productHandler := handlers.NewProductHandler()

	// API 路由组
	api := r.Group("/api")
	{
		// 商品相关路由
		products := api.Group("/products")
		{
			products.POST("", productHandler.CreateProduct)           // 创建商品
			products.GET("", productHandler.GetProducts)             // 获取商品列表
			products.GET("/:id", productHandler.GetProduct)          // 获取商品详情
			products.PUT("/:id", productHandler.UpdateProduct)       // 更新商品
			products.DELETE("/:id", productHandler.DeleteProduct)    // 删除商品
		}
	}

	// Swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
