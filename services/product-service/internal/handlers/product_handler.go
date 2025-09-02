package handlers

import (
	"net/http"
	"strconv"
	"oldoldcoder.com/cloudshop/product-service/internal/models"
	"oldoldcoder.com/cloudshop/product-service/internal/services"

	"github.com/gin-gonic/gin"
)

// ProductHandler 商品处理器
type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler 创建商品处理器
func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		productService: services.NewProductService(),
	}
}

// CreateProduct 创建商品
// @Summary 创建商品
// @Description 创建新商品
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param product body models.Product true "商品信息"
// @Success 200 {object} Response
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误",
			Data:    nil,
		})
		return
	}

	if err := h.productService.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "创建商品失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "创建商品成功",
		Data:    product,
	})
}

// GetProduct 获取商品详情
// @Summary 获取商品详情
// @Description 根据ID获取商品详情
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} Response
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "商品ID格式错误",
			Data:    nil,
		})
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "商品不存在",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取商品成功",
		Data:    product,
	})
}

// GetProducts 获取商品列表
// @Summary 获取商品列表
// @Description 分页获取商品列表
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param category_id query int false "分类ID"
// @Success 200 {object} Response
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	categoryIDStr := c.Query("category_id")

	var categoryID *uint
	if categoryIDStr != "" {
		if id, err := strconv.ParseUint(categoryIDStr, 10, 32); err == nil {
			temp := uint(id)
			categoryID = &temp
		}
	}

	products, total, err := h.productService.GetProducts(page, pageSize, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "获取商品列表失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取商品列表成功",
		Data: gin.H{
			"products": products,
			"total":    total,
			"page":     page,
			"page_size": pageSize,
		},
	})
}

// UpdateProduct 更新商品
// @Summary 更新商品
// @Description 更新商品信息
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Param product body models.Product true "商品信息"
// @Success 200 {object} Response
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "商品ID格式错误",
			Data:    nil,
		})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误",
			Data:    nil,
		})
		return
	}

	product.ID = uint(id)
	if err := h.productService.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "更新商品失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "更新商品成功",
		Data:    product,
	})
}

// DeleteProduct 删除商品
// @Summary 删除商品
// @Description 删除指定商品
// @Tags 商品管理
// @Accept json
// @Produce json
// @Param id path int true "商品ID"
// @Success 200 {object} Response
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "商品ID格式错误",
			Data:    nil,
		})
		return
	}

	if err := h.productService.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "删除商品失败",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "删除商品成功",
		Data:    nil,
	})
}

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
