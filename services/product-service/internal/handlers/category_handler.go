package handlers

import (
	"net/http"
	"strconv"

	"oldoldcoder.com/cloudshop/product-service/internal/models"
	"oldoldcoder.com/cloudshop/product-service/internal/services"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct{ svc *services.CategoryService }

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{svc: services.NewCategoryService()}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var m models.Category
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

func (h *CategoryHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	m, err := h.svc.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, Response{Code: 404, Message: "不存在"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK", Data: m})
}

func (h *CategoryHandler) List(c *gin.Context) {
	list, err := h.svc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "查询失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK", Data: list})
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m models.Category
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: 400, Message: "参数错误"})
		return
	}
	m.ID = uint(id)
	if err := h.svc.Update(&m); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "更新失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK", Data: m})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.svc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "删除失败"})
		return
	}
	c.JSON(http.StatusOK, Response{Code: 200, Message: "OK"})
}
