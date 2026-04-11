package handlers

import (
	"net/http"

	"github.com/IgorShirokov/marketplace/internal/catalog/application/queries"
	"github.com/gin-gonic/gin"
)

type CatalogItemsHandler struct {
	catalogItems *queries.CatalogItemsHandler
}

func NewCatalogItemsHandler(catalogItems *queries.CatalogItemsHandler) *CatalogItemsHandler {
	return &CatalogItemsHandler{catalogItems: catalogItems}
}

func (h *CatalogItemsHandler) GetCatalogItems(c *gin.Context) {
	items, err := h.catalogItems.Handle(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"catalogItems": items})
}
