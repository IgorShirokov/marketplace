package handlers

import (
	"net/http"

	"github.com/IgorShirokov/marketplace/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type CategoriesHandler struct {
	categories *queries.CategoriesHandler
}

func NewCategoryHandler(categories *queries.CategoriesHandler) *CategoriesHandler {
	return &CategoriesHandler{categories: categories}
}

func (h *CategoriesHandler) GetCategories(c *gin.Context) {
	categories, err := h.categories.Handle(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}
