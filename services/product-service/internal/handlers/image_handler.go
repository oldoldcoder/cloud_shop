package handlers

import (
	"net/http"
	"strconv"

	"oldoldcoder.com/cloudshop/product-service/internal/models"
	"oldoldcoder.com/cloudshop/product-service/internal/services"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct{ svc *services.ImageService }

func NewImageHandler() *ImageHandler { return &ImageHandler{svc: services.NewImageService()} }

func (h *ImageHandler) Create(c *gin.Context) {
	var m models.ProductImage
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误"})
		return
	}
	if err := h.svc.Create(&m); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "创建失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK", Data: m})
}

func (h *ImageHandler) ListByProduct(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("id"))
	list, err := h.svc.ListByProduct(uint(pid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK", Data: list})
}

func (h *ImageHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "删除失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK"})
}
