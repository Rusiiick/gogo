package api

import (
	"candy_shop/category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSup(c *gin.Context) {
	var supply SupplyDTO

	err := c.BindJSON(&supply)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.CreateSuppliy(c.Request.Context(), DTOtoModels(supply))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetSupply(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	supply, err := h.db.GetSupplyByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if supply == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, ToDTO(supply))
}

func (h *Handler) UpdateSup(c *gin.Context) {
	var supply SupplyDTO

	err := c.BindJSON(&supply)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	err = h.db.UpdateSupply(c.Request.Context(), DTOtoModels(supply), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении товара"})
		return
	}

	sup, err := h.db.GetSupplyByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении товара"})
		return
	}
	if sup == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Товар не найден"})
		return
	}

	c.JSON(http.StatusOK, ToDTO(sup))
}

func (h *Handler) DeleteSup(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	err = h.db.DeleteSupplyByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении задачи"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) CreateCtgry(c *gin.Context) {
	var category *CategoryDTO

	err := c.BindJSON(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.CreateCategory(c.Request.Context(), DTOtoModelsctgry(category))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetCtgryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	category, err := h.db.GetCategoryById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, ToDTOctgry(category))
}

func (h *Handler) GetAllCtgry(c *gin.Context) {
	category, err := h.db.GetAllCategory(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Категория не найдена"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) GetSupplyByCtgry(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	supplies, err := h.db.GetSupplyByCategoryID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if supplies == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Товары не найдены"})
		return
	}

	c.JSON(http.StatusOK, supplies)
}

func (h *Handler) UpdateCtgry(c *gin.Context) {
	var categories CategoryDTO

	err := c.BindJSON(&categories)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	err = h.db.UpdateCategory(c.Request.Context(), *DTOtoModelsctgry(&categories), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении категории"})
		return
	}

	ctgr, err := h.db.GetCategoryById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении товара"})
		return
	}
	if ctgr == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Товар не найден"})
		return
	}

	c.JSON(http.StatusOK, ToDTOctgry(ctgr))
}

func (h *Handler) DeleteCtgry(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	err = h.db.DeleteCategory(c.Request.Context(), category.Category{}, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении категории"})
		return
	}

	c.Status(http.StatusOK)
}
