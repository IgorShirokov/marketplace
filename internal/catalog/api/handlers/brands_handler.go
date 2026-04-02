package handlers

import (
	"net/http"

	"github.com/IgorShirokov/marketplace/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type BrandsHandler struct {
	brands *queries.BrandsHandler
}

func NewBrandsHandler(brands *queries.BrandsHandler) *BrandsHandler {
	return &BrandsHandler{brands: brands}
}

func (h *BrandsHandler) GetBrands(c *gin.Context) {
	brands, err := h.brands.Handle(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"brands": brands})
}
