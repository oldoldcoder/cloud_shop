package handlers

import (
	"net/http"
	"strconv"

	"oldoldcoder.com/cloudshop/product-service/internal/models"
	"oldoldcoder.com/cloudshop/product-service/internal/services"

	"github.com/gin-gonic/gin"
)

type StockHandler struct{ svc *services.StockService }

func NewStockHandler() *StockHandler { return &StockHandler{svc: services.NewStockService()} }

func (h *StockHandler) Get(c *gin.Context) {
	sku, _ := strconv.Atoi(c.Param("sku_id"))
	m, err := h.svc.GetBySKU(uint(sku))
	if err != nil {
		c.JSON(http.StatusNotFound, Response{Code: 404, Message: "不存在"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK", Data: m})
}

func (h *StockHandler) Upsert(c *gin.Context) {
	var m models.ProductStock
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误"})
		return
	}
	if err := h.svc.Upsert(&m); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "保存失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK", Data: m})
}

func (h *StockHandler) Increase(c *gin.Context) {
	sku, _ := strconv.Atoi(c.Param("sku_id"))
	delta, _ := strconv.Atoi(c.Query("delta"))
	if err := h.svc.Increase(uint(sku), delta); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "增加失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK"})
}

func (h *StockHandler) Decrease(c *gin.Context) {
	sku, _ := strconv.Atoi(c.Param("sku_id"))
	delta, _ := strconv.Atoi(c.Query("delta"))
	if err := h.svc.Decrease(uint(sku), delta); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "减少失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK"})
}
